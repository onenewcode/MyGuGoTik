package main

import (
	"GuGoTik/src/constant/config"
	"GuGoTik/src/extra/profiling"
	"GuGoTik/src/extra/tracing"
	"GuGoTik/src/models"
	"GuGoTik/src/rpc/auth"
	"GuGoTik/src/storage/database"
	"GuGoTik/src/storage/redis"
	"GuGoTik/src/utils/consul"
	"GuGoTik/src/utils/logging"
	"GuGoTik/src/utils/prom"
	"context"
	"github.com/bits-and-blooms/bloom/v3"
	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	redis2 "github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"net/http"
	"os"
	"syscall"
)

func main() {
	tp, err := tracing.SetTraceProvider(config.AuthRpcServerName)

	if err != nil {
		// 方法添加了一个字段到日志条目中
		logging.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Panicf("Error to set the trace")
	}
	defer func() {
		// 等待所有追踪数据被导出完成，然后释放资源
		if err := tp.Shutdown(context.Background()); err != nil {
			logging.Logger.WithFields(logrus.Fields{
				"err": err,
			}).Errorf("Error to set the trace")
		}
	}()

	// Configure Pyroscope 用于服务器信息收集
	profiling.InitPyroscope("GuGoTik.AuthService")
	// 添加服务字段
	log := logging.LogService(config.AuthRpcServerName)
	// 开启rpc 服务监听
	lis, err := net.Listen("tcp", config.EnvCfg.PodIpAddr+config.AuthRpcServerPort)

	if err != nil {
		log.Panicf("Rpc %s listen happens error: %v", config.AuthRpcServerName, err)
	}
	// 创建一个新的rpc监控指标收集器
	srvMetrics := grpcprom.NewServerMetrics(
		// 每个bucket代表一个区间，用于累计落在该区间内的请求处理时间
		grpcprom.WithServerHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
	)

	reg := prom.Client
	// 监控指标注入到 prom的服务器
	reg.MustRegister(srvMetrics)
	//创建一个目标误报率为 0.1% 的新 Bloom 过滤器
	// Create a new Bloom filter with a target false positive rate of 0.1%
	BloomFilter = bloom.NewWithEstimates(10000000, 0.001) // assuming we have 1 million users

	// Initialize BloomFilter from database
	var users []models.User
	userNamesResult := database.Client.WithContext(context.Background()).Select("user_name").Find(&users)
	if userNamesResult.Error != nil {
		log.Panicf("Getting user names from databse happens error: %s", userNamesResult.Error)
		panic(userNamesResult.Error)
	}
	for _, u := range users {
		BloomFilter.AddString(u.UserName)
	}

	// Create a go routine to receive redis message and add it to BloomFilter
	go func() {
		// 利用发布订阅模型上传用户数据
		pubSub := redis.Client.Subscribe(context.Background(), config.BloomRedisChannel)
		defer func(pubSub *redis2.PubSub) {
			err := pubSub.Close()
			if err != nil {
				log.Panicf("Closing redis pubsub happend error: %s", err)
			}
		}(pubSub)

		_, err := pubSub.ReceiveMessage(context.Background())
		if err != nil {
			log.Panicf("Reveiving message from redis happens error: %s", err)
			panic(err)
		}

		ch := pubSub.Channel()
		for msg := range ch {
			log.Infof("Add user name to BloomFilter: %s", msg.Payload)
			BloomFilter.AddString(msg.Payload)
		}
	}()
	// 新建一个rpc服务
	s := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		grpc.ChainUnaryInterceptor(srvMetrics.UnaryServerInterceptor(grpcprom.WithExemplarFromContext(prom.ExtractContext))), //用于信息监控
		grpc.ChainStreamInterceptor(srvMetrics.StreamServerInterceptor(grpcprom.WithExemplarFromContext(prom.ExtractContext))),
	)
	// consul注册服务
	if err := consul.RegisterConsul(config.AuthRpcServerName, config.AuthRpcServerPort); err != nil {
		log.Panicf("Rpc %s register consul happens error for: %v", config.AuthRpcServerName, err)
	}
	log.Infof("Rpc %s is running at %s now", config.AuthRpcServerName, config.AuthRpcServerPort)

	var srv AuthServiceImpl
	// rpc服务进行在注册
	auth.RegisterAuthServiceServer(s, srv)
	// rpc health注册
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	srv.New()
	srvMetrics.InitializeMetrics(s)
	// 并发地管理和协调多个任务
	g := &run.Group{}
	// 启动rpc服务
	g.Add(func() error {
		return s.Serve(lis)
	}, func(err error) {
		s.GracefulStop()
		s.Stop()
		log.Errorf("Rpc %s listen happens error for: %v", config.AuthRpcServerName, err)
	})
	// 创建HTTP服务器实例，该服务用于向pro提供检测指标
	httpSrv := &http.Server{Addr: config.EnvCfg.PodIpAddr + config.Metrics}
	g.Add(func() error {
		m := http.NewServeMux()
		m.Handle("/metrics", promhttp.HandlerFor(
			reg,
			promhttp.HandlerOpts{
				EnableOpenMetrics: true,
			},
		))
		httpSrv.Handler = m
		log.Infof("Promethus now running")
		return httpSrv.ListenAndServe()
	}, func(error) {
		if err := httpSrv.Close(); err != nil {
			log.Errorf("Prometheus %s listen happens error for: %v", config.AuthRpcServerName, err)
		}
	})
	//一个信号处理器，用于监听两种系统信号：SIGINT（通常由用户按下Ctrl+C触发，意在请求进程中断）和SIGTERM（一个典型的进程终止信号，通常由系统或管理员发出，要求进程正常退出）
	g.Add(run.SignalHandler(context.Background(), syscall.SIGINT, syscall.SIGTERM))

	if err := g.Run(); err != nil {
		log.WithFields(logrus.Fields{
			"err": err,
		}).Errorf("Error when runing http server")
		os.Exit(1)
	}
}
