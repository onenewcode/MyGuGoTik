package profiling

import (
	"GuGoTik/src/constant/config"
	"GuGoTik/src/utils/logging"
	"github.com/grafana/pyroscope-go"
	log "github.com/sirupsen/logrus"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"os"
	"runtime"
)

// 初始化并启动Pyroscope性能监控工具。Pyroscope是一个可观察性工具，用于监控和分析Go应用的性能，特别是CPU和内存使用情况。下面是代码的逐行解释：
func InitPyroscope(appName string) {
	// 判断是否开启性能监控
	if config.EnvCfg.PyroscopeState != "enable" {
		logging.Logger.WithFields(log.Fields{
			"appName": appName,
		}).Infof("User close Pyroscope, the service would not run.")
		return
	}
	// 设置了互斥锁分析的样本率
	runtime.SetMutexProfileFraction(5)
	// 阻塞事件分析的样本率
	runtime.SetBlockProfileRate(5)

	_, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: appName,
		ServerAddress:   config.EnvCfg.PyroscopeAddr,
		Logger:          logrus.NewWriter(),
		Tags:            map[string]string{"hostname": os.Getenv("HOSTNAME")},
		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	})

	if err != nil {
		logging.Logger.WithFields(log.Fields{
			"appName": appName,
			"err":     err,
		}).Warnf("Pyroscope failed to run.")
		return
	}
}
