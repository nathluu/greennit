package main

import (
	"flag"

	"github.com/nathluu/greennit/internal/app/api"
	"github.com/nathluu/greennit/internal/pkg/config"
	"github.com/nathluu/greennit/internal/pkg/http/server"
	"github.com/nathluu/greennit/internal/pkg/log"
	"github.com/spf13/viper"
)

func main() {
	confPath := flag.String("confPath", "configs", "Configuration path")
	confName := flag.String("confName", "greennit", "Configuration file")
	flag.Parse()

	log.Init(log.Fields{
		"service": "greennit",
	})

	log.Debugf("ConfiPath: %s\n", *confPath)
	log.Debugf("ConfigName: %s\n", *confName)

	conf := config.Config{
		Viper:      viper.New(),
		ConfigPath: *confPath,
		ConfigName: *confName,
	}

	if err:= conf.Init(); err != nil {
		log.Panicf("cannot read configuration: %v\n", err)
	}

	router, err := api.NewRouter(&conf)
	if err != nil {
		log.Panicf("failed to init router: %v\n", err)
	}
	log.Debugf("server configuration: %v\n", conf.HTTPServer)
	log.Debugf("server configuration: %T\n", conf.HTTPServer.ReadHeaderTimeout)
	server.ListenAndServe(conf.HTTPServer, router)
}
