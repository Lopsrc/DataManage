package suite

import (
	"context"
	"net"
	"testing"

	manager1 "server/protos/gen/go/manager"
	"server/server/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
const (
	grpsHost = "localhost"
)
type Suite struct{
	*testing.T
	Cfg *config.Config
	Price manager1.ManagerPriceClient
	Work manager1.ManagerWorkClient
}

func NewSuit(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	cfg := config.GetConfig("../../config/local_test.yaml")
	ctx, cancelCtx := context.WithTimeout(context.Background(), cfg.GRPC.Timeout)

	t.Cleanup(func(){
		t.Helper()
		cancelCtx()
	})

	cc, err := grpc.DialContext(context.Background(), grpcAddress(cfg), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("grpc server connection failed: %v", err)
	}

	return ctx, &Suite{
		T: t,
		Cfg: cfg,
		Work: manager1.NewManagerWorkClient(cc),
		Price: manager1.NewManagerPriceClient(cc),
	}
}

func grpcAddress(cfg *config.Config) string{
	return net.JoinHostPort(grpsHost, cfg.GRPC.Port)
}