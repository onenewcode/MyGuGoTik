package logging

import (
	"GuGoTik/src/constant/config"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"io"
	"os"
	"path"
)

var hostname string

func init() {
	hostname, _ = os.Hostname()

	switch config.EnvCfg.LoggerLevel {
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "WARN", "WARNING":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "FATAL":
		log.SetLevel(log.FatalLevel)
	case "TRACE":
		log.SetLevel(log.TraceLevel)
	}
	// 设置文件输出路径
	filePath := path.Join("/var", "log", "gugotik", "gugotik.log")

	dir := path.Dir(filePath)
	if err := os.MkdirAll(dir, os.FileMode(0755)); err != nil {
		panic(err)
	}

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	// 设置输出格式
	log.SetFormatter(&log.JSONFormatter{})
	// 添加回调函数
	log.AddHook(logTraceHook{})
	// 设置多重输出，同时输出到命令行和指定文件
	log.SetOutput(io.MultiWriter(f, os.Stdout))

	Logger = log.WithFields(log.Fields{
		"Tied":     config.EnvCfg.TiedLogging,
		"Hostname": hostname,
		"PodIP":    config.EnvCfg.PodIpAddr,
	})
}

type logTraceHook struct{}

func (t logTraceHook) Levels() []log.Level { return log.AllLevels }

// 用于集成OpenTelemetry的分布式追踪功能到logrus日志库中
func (t logTraceHook) Fire(entry *log.Entry) error {
	ctx := entry.Context
	if ctx == nil {
		return nil
	}

	span := trace.SpanFromContext(ctx)
	sCtx := span.SpanContext()
	if sCtx.HasTraceID() {
		entry.Data["trace_id"] = sCtx.TraceID().String()
	}
	if sCtx.HasSpanID() {
		entry.Data["span_id"] = sCtx.SpanID().String()
	}

	if config.EnvCfg.LoggerWithTraceState == "enable" { //判断是否开启日志追踪
		attrs := make([]attribute.KeyValue, 0)
		// 存储日志的严重性级别
		logSeverityKey := attribute.Key("log.severity")
		//存储日志消息的内容
		logMessageKey := attribute.Key("log.message")
		attrs = append(attrs, logSeverityKey.String(entry.Level.String()))
		attrs = append(attrs, logMessageKey.String(entry.Message))
		for key, value := range entry.Data {
			fields := attribute.Key(fmt.Sprintf("log.fields.%s", key))
			attrs = append(attrs, fields.String(fmt.Sprintf("%v", value)))
		}
		//在当前的追踪Span中添加一个名为“log”的事件，通过trace.WithAttributes(attrs...)将前面收集的所有属性传递给这个事件。
		//这样，日志的详细信息就会和这个追踪事件关联起来，便于在追踪结果中查看和分析
		span.AddEvent("log", trace.WithAttributes(attrs...))
		// 判断日志的错我是否等级太大
		if entry.Level <= log.ErrorLevel {
			span.SetStatus(codes.Error, entry.Message)
		}
	}
	return nil
}

var Logger *log.Entry

func LogService(name string) *log.Entry {
	return Logger.WithFields(log.Fields{
		"Service": name,
	})
}

func SetSpanError(span trace.Span, err error) {
	span.RecordError(err)
	span.SetStatus(codes.Error, "Internal Error")
}

func SetSpanErrorWithDesc(span trace.Span, err error, desc string) {
	span.RecordError(err)
	span.SetStatus(codes.Error, desc)
}

func SetSpanWithHostname(span trace.Span) {
	span.SetAttributes(attribute.String("hostname", hostname))
	span.SetAttributes(attribute.String("podIP", config.EnvCfg.PodIpAddr))
}
