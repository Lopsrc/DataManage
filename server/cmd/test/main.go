package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	grpc_ "server/server/internal/app"
	"server/server/internal/config"
	"server/server/pkg/client/postgresql"
	"syscall"
)

const (
	pathConfig = "server/config/local_test.yaml"
)

func main() {
	
	cfg := config.GetConfig(pathConfig)

	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	postgreSQLClient, err := postgresql.NewClient(context.Background(), 3, cfg.Storage)
	if err!= nil {
        panic(err)
    }

	application := grpc_.New(log, cfg.GRPC.Port, postgreSQLClient)

	go func() {
		application.GRPCServer.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCServer.Stop()
	log.Info("gRPC server stopped")
}