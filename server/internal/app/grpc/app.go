package grpc_

import (
	"fmt"
	"log/slog"
	"net"
	
	price "server/server/internal/grpc/manager/price"
	work "server/server/internal/grpc/manager/work"

	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       string
}

func New(log *slog.Logger, w work.WorkService, p price.PriceService , port string) *App {
	gRPCServer := grpc.NewServer()
	work.Register(gRPCServer, w)
	price.Register(gRPCServer, p)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}
// Run gRPC server.
func (app *App) MustRun() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func (app *App) Run() error {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", app.port))
	if err != nil {
		return fmt.Errorf("Run(): %w", err)
	}

	app.log.Info("grpc server started", slog.Attr{Key: "address: ", Value: slog.StringValue(listener.Addr().String())})

	if err := app.gRPCServer.Serve(listener); err != nil {
		return fmt.Errorf("Run(): %w", err)
	}
	return nil
}

func (app *App) Stop(){
	app.log.Info("stopping gRPC server", slog.String("port", app.port))
	app.gRPCServer.GracefulStop()
}