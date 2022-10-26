package main

import (
	"flag"
	"fmt"
	"github.com/charlie-bit/yanxue/config"
	"github.com/charlie-bit/yanxue/controller/api"
	"github.com/charlie-bit/yanxue/db"
	"github.com/charlie-bit/yanxue/log"
	"github.com/charlie-bit/yanxue/uhttp"
	"os"
	"os/signal"
	"syscall"
)

var (
	Version = "no version provide"
)

// @title Mall Blockchain API Service
// @version 1.0
// @Produce json
// @contact.name API Support
// @license.name Apache 2.0
func main() {
	var (
		configFile  string
		ping        bool
		showVersion bool
	)

	flag.StringVar(&configFile, "f", "config/setting.yml", "config file")
	flag.BoolVar(&ping, "ping", false, "Ping server.")
	flag.BoolVar(&showVersion, "version", false, "Print version information.")
	flag.BoolVar(&showVersion, "v", false, "Print version information.")
	flag.Parse()

	// get config
	cfg, err := config.NewConfig(configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
		return
	}

	fmt.Printf("version: %s\n", Version)
	fmt.Printf("env: %s\n", cfg.Env)

	log.Init()

	// test network
	if ping {
		if err = cfg.Ping(); err != nil {
			log.Info(err.Error())
		} else {
			log.Info("hello")
		}
		return
	}

	if err = db.InitMysql(cfg.Env); err != nil {
		log.Error(err.Error())
		return
	}

	if err = db.InitRedis(cfg); err != nil {
		log.Error(err.Error())
		return
	}

	_ = uhttp.GetHTTPClient()

	// start api server
	api.Start()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit

	log.Info("evm chain api exit success")
	log.Exit()
}
