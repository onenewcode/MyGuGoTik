package rabbitmq

import (
	"context"
	"go.opentelemetry.io/otel"
)

type AmqpHeadersCarrier map[string]interface{}

func (a AmqpHeadersCarrier) Get(key string) string {
	v, ok := a[key]
	if !ok {
		return ""
	}
	return v.(string)
}

func (a AmqpHeadersCarrier) Set(key string, value string) {
	a[key] = value
}

func (a AmqpHeadersCarrier) Keys() []string {
	i := 0
	r := make([]string, len(a))

	for k := range a {
		r[i] = k
		i++
	}

	return r
}

// 初始化分布式追踪
// 这个返回的字典可以直接用于AMQP消息的头部，从而在消息传递过程中传播追踪数据，使得消息在不同服务间流转时，依然能维持分布式追踪链路的连续性
func InjectAMQPHeaders(ctx context.Context) map[string]interface{} {
	h := make(AmqpHeadersCarrier)
	//取全局的文本映射传播器,
	otel.GetTextMapPropagator().Inject(ctx, h)
	return h
}

func ExtractAMQPHeaders(ctx context.Context, headers map[string]interface{}) context.Context {
	return otel.GetTextMapPropagator().Extract(ctx, AmqpHeadersCarrier(headers))
}
