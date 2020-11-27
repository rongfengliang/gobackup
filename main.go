package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/huacnlee/gobackup/config"
	"github.com/huacnlee/gobackup/logger"
	"github.com/huacnlee/gobackup/model"
	"github.com/robfig/cron"
	"gopkg.in/urfave/cli.v1"
)

const (
	usage = "Easy full stack backup operations on UNIX-like systems"
)

var (
	modelName = ""
	version   = "master"
)

func main() {
	app := cli.NewApp()
	app.Version = version
	app.Name = "gobackup"
	app.Usage = usage

	app.Commands = []cli.Command{
		{
			Name: "perform",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "model, m",
					Usage:       "Model name that you want execute",
					Destination: &modelName,
				},
			},
			Action: func(c *cli.Context) error {
				if len(modelName) == 0 {
					performAll()
				} else {
					performOne(modelName)
				}

				return nil
			},
		},
		{
			Name: "start",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "start",
					Usage: "cron  execute backup task",
				},
			},
			Action: func(c *cli.Context) error {
				cronPerformAll()
				sigChan := make(chan os.Signal, 1)
				signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
				sig := <-sigChan
				logger.Info("shutting down %v signal received", sig)
				return nil
			},
		},
	}
	app.Run(os.Args)
}

func performAll() {
	for _, modelConfig := range config.Models {
		m := model.Model{
			Config: modelConfig,
		}
		m.Perform()
	}
}

func performOne(modelName string) {
	for _, modelConfig := range config.Models {
		if modelConfig.Name == modelName {
			m := model.Model{
				Config: modelConfig,
			}
			m.Perform()
			return
		}
	}
}

func cronPerformAll() {
	logger.Info(config.Scheduler)
	mycron := cron.New()
	if err := mycron.AddFunc(config.Scheduler, func() {
		for _, modelConfig := range config.Models {
			m := model.Model{
				Config: modelConfig,
			}
			m.Perform()
		}
	}); err != nil {
		logger.Error("start cron service error:", err.Error())
		os.Exit(1)
	}
	mycron.Start()
}
