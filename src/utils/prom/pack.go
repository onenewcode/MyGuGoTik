package prom

import "github.com/prometheus/client_golang/prometheus"

// 注册表
var Client = prometheus.NewRegistry()
