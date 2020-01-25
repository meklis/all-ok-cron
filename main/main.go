package main

import (
	"flag"
	"github.com/meklis/all-ok-cron/config"
	"github.com/meklis/all-ok-cron/cron"
	"github.com/meklis/http-snmpwalk-proxy/logger"
	"os"
)

var (
	Config     config.Configuration
	pathConfig string
	lg         *logger.Logger
)

func init() {
	flag.StringVar(&pathConfig, "c", "crontab.conf.yml", "Configuration file")
	flag.Parse()
}

func main() {
	if err := config.LoadConfig(pathConfig, &Config); err != nil {
		panic(err)
	}
	if Config.Logger.Console.Enabled {
		color := 0
		if Config.Logger.Console.EnabledColor {
			color = 1
		}
		lg, _ = logger.New("pooler", color, os.Stdout)
		lg.SetLogLevel(logger.LogLevel(Config.Logger.Console.LogLevel))
		if !Config.Logger.Console.PrintDebugLine {
			lg.SetFormat("#%{id} %{time} > %{level} %{message}")
		} else {
			lg.SetFormat("#%{id} %{time} (%{filename}:%{line}) > %{level} %{message}")
		}
	} else {
		lg, _ = logger.New("no_log", 0, os.DevNull)
	}

	sh := cron.Init(Config.Crontab, lg)
	sh.Run()
}
