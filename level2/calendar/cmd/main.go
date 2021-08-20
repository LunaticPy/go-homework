package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"task/level2/calendar/config"
	"task/level2/calendar/internal/server"

	"github.com/sirupsen/logrus"
)

var configPath = "./config/config.json"

func init() {
	flag.StringVar(&configPath, "config-path", "config/config.json", "path to config file")
}

func main() {
	flag.Parse()
	cfg, err := config.GetConfig(configPath)
	osSignals := make(chan os.Signal, 1)
	listenErr := make(chan error, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)
	if err != nil {
		logrus.Error(err)
		return
	}
	serv := server.NewServer(cfg)

	serv.Start(&listenErr)
	select {
	case err := <-listenErr:
		if err != nil {
			logrus.Error(err)
			os.Exit(1)
		}
	case <-osSignals:
		serv.Stop()
	}
}
