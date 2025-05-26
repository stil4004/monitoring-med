package config

import (
	"service/pkg/secure"

	"go.uber.org/zap"
	//"market_api/pkg/secure"
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	LoggerZapConfig zap.Config
	Manager         ManagerStatConfig
	InnerConnection InnerConnectionConfig
	Deploy          Deploy
	Server          ServerConfig
	Jaeger          JaegerConfig
	VictoriaMetrics VictoriaMetricsConfig
}

type VictoriaMetricsConfig struct {
	BaseURL string `json:"baseURL"`
}

type JaegerConfig struct {
	Host        string `json:"host"`
	Port        string `json:"port"`
	ServiceName string `json:"serviceName"`
}

type ServerConfig struct {
	AppVersion string `json:"appVersion"`
	Host       string `json:"host" validate:"required"`
	Port       string `json:"port" validate:"required"`
	Timeout    time.Duration
}

type Deploy struct {
	IsTest bool   `json:"is_test"`
	Port   string `json:"port"`
}

type InnerConnectionConfig struct {
	ReloadTime time.Duration `json:"ReloadTime"`
}

//type PrometheusConfig struct {
//	Host string `json:"host"`
//	Port string `json:"port"`
//}

type ManagerStatConfig struct {
	FlushInterval time.Duration `json:"flushInterval"`
}

type LoggerConfig struct {
	Level string `json:"level"`
}

func LoadConfig(isTest bool) (*viper.Viper, error) {

	var configName string
	viperInstance := viper.New()

	if isTest {
		configName = "config_test"
	} else {
		configName = "config"
	}

	viperInstance.AddConfigPath("./config")
	viperInstance.SetConfigName(configName)
	viperInstance.SetConfigType("yml")

	err := viperInstance.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return viperInstance, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {

	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode config into struct, %v", err)
		return nil, err
	}
	return &c, nil
}

func DecryptConfig(cfg *Config, s *secure.Shield) {

	// cfg.Postgres.Host = s.DecryptMessage(cfg.Postgres.Host)
	// cfg.Postgres.Port = s.DecryptMessage(cfg.Postgres.Port)
	// cfg.Postgres.User = s.DecryptMessage(cfg.Postgres.User)
	// cfg.Postgres.Password = s.DecryptMessage(cfg.Postgres.Password)
	// cfg.Postgres.DBName = s.DecryptMessage(cfg.Postgres.DBName)

	//cfg.ElcConfig.Address = s.DecryptMessage(cfg.ElcConfig.Address)
	//cfg.ElcConfig.Type = s.DecryptMessage(cfg.ElcConfig.Type)

	//cfg.Qca.ApiUrl = s.DecryptMessage(cfg.Qca.ApiUrl)
	//cfg.Qca.ApiKey = s.DecryptMessage(cfg.Qca.ApiKey)

	// cfg.Jaeger.ServiceName = s.DecryptMessage(cfg.Jaeger.ServiceName)
	// cfg.Jaeger.LocalAgentHostPort = s.DecryptMessage(cfg.Jaeger.LocalAgentHostPort)

	// cfg.Nats.Host = s.DecryptMessage(cfg.Nats.Host)
	// cfg.Nats.Port = s.DecryptMessage(cfg.Nats.Port)
	// cfg.Nats.Client = s.DecryptMessage(cfg.Nats.Client)
}
