package config

import (
	"github.com/nathluu/greennit/internal/pkg/db/mongodb"
	"github.com/nathluu/greennit/internal/pkg/http/server"
	"github.com/nathluu/greennit/internal/pkg/log"
	"github.com/spf13/viper"
)

type (
	Config struct {
		Viper      *viper.Viper
		ConfigPath string
		ConfigName string
		//LogLevel   string         `mapstructure:"log_level"`
		HTTPServer server.Config  `mapstructure:"http_server"`
		Database   mongodb.Config `mapstructure:"mongodb"`
	}
)

func (conf *Config) Init() error {
	conf.Viper.AddConfigPath(conf.ConfigPath)
	conf.Viper.SetConfigName(conf.ConfigName)
	if err := conf.Viper.ReadInConfig(); err != nil {
		log.Errorf("cannot read config from file: %v", err)
		return err
	}

	if err := conf.Viper.Unmarshal(&conf); err != nil {
		log.Errorf("unmarshal configuration failed: %v\n", err)
		return err
	}

	return nil
}