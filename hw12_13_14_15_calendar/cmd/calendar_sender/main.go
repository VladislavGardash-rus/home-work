package main

import (
	"context"
	"flag"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/cfg"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/cmd"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/logger"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/services"
	"os"
	"os/signal"
	"syscall"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "config.json", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		cmd.PrintVersion("0.0.1", "01.02.2024", "")
		os.Exit(0)
	}

	err := cfg.InitConfig(configFile)
	if err != nil {
		panic(err)
	}

	err = logger.InitLogger(cfg.Config().Logger.Level)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	go watchExitSignals(cancel)

	eventSenderService := services.NewEventSenderService()
	go eventSenderService.Start(ctx)

	logger.UseLogger().Info("calendar_sender service is running...")

	<-ctx.Done()

	logger.UseLogger().Info("calendar_sender service was stopped")
}

func watchExitSignals(cancel context.CancelFunc) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	cancel()
}
