package tracing

import (
	"GuGoTik/src/constant/config"
	"GuGoTik/src/utils/logging"
	"context"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
	trace2 "go.opentelemetry.io/otel/trace"
)

var Tracer trace2.Tracer

// 初始化一个Trace的提供者
func SetTraceProvider(name string) (*trace.TracerProvider, error) {
	// 初始化一个http客户端，用于提交监控的信息
	client := otlptracehttp.NewClient(
		otlptracehttp.WithEndpoint(config.EnvCfg.TracingEndPoint), // 配置了追踪数据的接收端点
		otlptracehttp.WithInsecure(),                              //不使用TLS
	)
	// 创建OTLP导出器
	exporter, err := otlptrace.New(context.Background(), client)
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorf("Can not init otel !")
		return nil, err
	}
	//Sampler接口定义了采样决策逻辑，即决定某次追踪数据是否应该被记录
	var sampler trace.Sampler
	if config.EnvCfg.OtelState == "disable" {
		sampler = trace.NeverSample() // 返回一个不追踪的采样
	} else {
		sampler = trace.TraceIDRatioBased(config.EnvCfg.OtelSampler) // 基于追踪ID的比例采样策略将被采用
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter), // 指定数据应当如何被批量发送到后端
		trace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL, // 指定了资源属性的语义命名空间URL，这是一个固定的URL指向OpenTelemetry规范中的资源属性schema，确保属性的语义被标准化
				semconv.ServiceNameKey.String(name),
			),
		),
		// 这个参数设置了采样器（Sampler），决定哪些追踪应当被记录和发送出去，哪些应当被丢弃
		trace.WithSampler(sampler),
	)
	otel.SetTracerProvider(tp)
	//传播器的作用是在进程间传播追踪上下文（trace context），确保跨服务的请求能被正确关联和追踪
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	Tracer = otel.Tracer(name)
	return tp, nil
}
