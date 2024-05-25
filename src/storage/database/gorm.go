package database

import (
	"GuGoTik/src/constant/config"
	"GuGoTik/src/utils/logging"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"gorm.io/plugin/opentelemetry/tracing"
	"strings"
	"time"
)

var Client *gorm.DB

func init() {
	var err error

	gormLogrus := logging.GetGormLogger()

	var cfg gorm.Config
	if config.EnvCfg.PostgreSQLSchema == "" {
		cfg = gorm.Config{
			PrepareStmt: true,
			Logger:      gormLogrus,
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: config.EnvCfg.PostgreSQLSchema + "." + config.EnvCfg.PostgreSQLPrefix,
			},
		}
	} else {
		cfg = gorm.Config{
			PrepareStmt: true,
			Logger:      gormLogrus,
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: config.EnvCfg.PostgreSQLSchema + "." + config.EnvCfg.PostgreSQLPrefix,
			},
		}
	}
	// 建立数据库连接
	if Client, err = gorm.Open(
		postgres.Open(
			fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
				config.EnvCfg.PostgreSQLHost,
				config.EnvCfg.PostgreSQLUser,
				config.EnvCfg.PostgreSQLPassword,
				config.EnvCfg.PostgreSQLDataBase,
				config.EnvCfg.PostgreSQLPort)),
		&cfg,
	); err != nil {
		panic(err)
	}

	if config.EnvCfg.PostgreSQLReplicaState == "enable" {
		var replicas []gorm.Dialector
		for _, addr := range strings.Split(config.EnvCfg.PostgreSQLReplicaAddress, ",") {
			pair := strings.Split(addr, ":")
			if len(pair) != 2 {
				continue
			}

			replicas = append(replicas, postgres.Open(
				fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
					pair[0],
					config.EnvCfg.PostgreSQLReplicaUsername,
					config.EnvCfg.PostgreSQLReplicaPassword,
					config.EnvCfg.PostgreSQLDataBase,
					pair[1])))
		}
		// 使用读写分离
		err := Client.Use(dbresolver.Register(dbresolver.Config{
			Replicas: replicas,
			Policy:   dbresolver.RandomPolicy{},
		}))
		if err != nil {
			panic(err)
		}
	}

	sqlDB, err := Client.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(24 * time.Hour)
	sqlDB.SetConnMaxIdleTime(time.Hour)
	// 注册追踪中间件
	if err := Client.Use(tracing.NewPlugin()); err != nil {
		panic(err)
	}
	// TODO 设置提交地址

	//if err := Client.Use(tracing.NewPlugin(tracing.WithTracerProvider())); err != nil {
	//	panic(err)
	//}
}
