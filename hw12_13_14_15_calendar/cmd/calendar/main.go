package main

import (
	"context"
	"flag"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/cfg"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/cmd"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/logger"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/storage"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/transport/grpc_server"
	http_server "github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/transport/http"
	"net"
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

	iStorage, err := storage.NewStorage(ctx, cfg.Config().Storage.Type, cfg.Config().Storage.Connection)
	if err != nil {
		panic(err)
	}

	httpServer := http_server.NewServer(net.JoinHostPort(cfg.Config().CalendarHttpServer.Host, cfg.Config().CalendarHttpServer.Port), iStorage, "calendar_http")
	go func() {
		err := httpServer.Start()
		if err != nil {
			logger.UseLogger().Error(err)
			cancel()
		}
	}()

	grpcServer := grpc_server.NewServer(net.JoinHostPort(cfg.Config().CalendarGrpcServer.Host, cfg.Config().CalendarGrpcServer.Port), iStorage, "calendar_grpc")
	go func() {
		err := grpcServer.Start()
		if err != nil {
			logger.UseLogger().Error(err)
			cancel()
		}
	}()

	logger.UseLogger().Info("calendar service is running...")

	<-ctx.Done()
	shutDownServers(ctx, httpServer, grpcServer)

	logger.UseLogger().Info("calendar service was stopped")
}

func watchExitSignals(cancel context.CancelFunc) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	cancel()
}

func shutDownServers(ctx context.Context, httpServer *http_server.Server, grpcServer *grpc_server.Server) {
	err := httpServer.Stop(ctx)
	if err != nil {
		logger.UseLogger().Error(err)
	}

	err = grpcServer.Stop(ctx)
	if err != nil {
		logger.UseLogger().Error(err)
	}
}
