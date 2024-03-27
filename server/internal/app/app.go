package app

import (
	"log/slog"

	grpc_ "server/server/internal/app/grpc"
	price "server/server/internal/services/price"
	work "server/server/internal/services/work"
	sqlwork "server/server/internal/storage/postgresql/manager/work"
	sqlprice "server/server/internal/storage/postgresql/manager/price"

	"github.com/jackc/pgx/v4/pgxpool"
)

type App struct{
	GRPCServer *grpc_.App
}

func New(log *slog.Logger, grpcPort string, pconn *pgxpool.Pool) *App{

	storageWork := sqlwork.New(pconn, log)
    storagePrice := sqlprice.New(pconn, log)

	priceService := price.New(storagePrice, log)
	workService := work.New(storageWork, log)

	grpcApp := grpc_.New(log, workService, priceService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}