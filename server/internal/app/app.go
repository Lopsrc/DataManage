package app

import (
	"log/slog"
	"time"

	grpc_ "server/server/internal/app/grpc"
	"server/server/internal/config"
	price "server/server/internal/services/price"
	work "server/server/internal/services/work"
)

type App struct{
	GRPCServer *grpc_.App
}

func New(log *slog.Logger, grpcPort string, cfg config.Config, token time.Duration) *App{

	//create pool connections.

	// create storage.
	
	priceService := price.New(storage, log)
	workService := work.New(storage, log)
	grpcApp := grpc_.New(log, workService, priceService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}