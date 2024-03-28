package price

import (
	"context"
	"errors"
	manager1 "server/protos/gen/go/manager"
	m "server/server/internal/middleware/manager/price"
	models "server/server/internal/models/price"
	price "server/server/internal/services/price"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
// go:generate go run github.com/vektra/mockery/v2@v2.42.0 --name=PriceService
type PriceService interface{
	Create( 
		ctx context.Context,
        rec models.CreatePrice,
	)error
	Update(
		ctx context.Context,
        rec models.UpdatePrice,
	)error
	Get(
		ctx context.Context,
        rec models.GetPrice,
	)(models.Prices, error)
}

type serverAPI struct {
	manager1.UnimplementedManagerPriceServer
	p PriceService
}

func Register(gRPC *grpc.Server, price PriceService) {
	manager1.RegisterManagerPriceServer(gRPC, &serverAPI{p: price})
}

func (s *serverAPI) Create(ctx context.Context, req *manager1.CreatePriceRequest) (*manager1.CreatePriceResponse, error){
	// Handle request
	err := m.HandleCreate(req)
	if err!= nil{
        return nil, status.Error(codes.InvalidArgument, err.Error())
    }
	// Create a record.
	err = s.p.Create(ctx, models.CreatePrice{
		ID: req.UserId,
        Price: req.Price,
    })
    if err!= nil { 
		if errors.Is(err, price.ErrAlreadyExists) {
            return nil, status.Error(codes.AlreadyExists, "entity already exists")
        }
        return nil, status.Error(codes.Internal, "internal error")
    }
	return &manager1.CreatePriceResponse{
		IsCreate: true,
	}, nil
}

func (s *serverAPI) Update(ctx context.Context, req *manager1.UpdatePriceRequest) (*manager1.UpdatePriceResponse, error){
	// Handle request
	err := m.HandleUpdate(req)
	if err!= nil{
        return nil, status.Error(codes.InvalidArgument, err.Error())
    }
	// Update a record.
	err = s.p.Update(ctx, models.UpdatePrice{
		ID: req.UserId,
        Price: req.Price,
    })
    if err!= nil {
		if errors.Is(err, price.ErrNotFound) {
            return nil, status.Error(codes.NotFound, "Entity not found")
        }
        return nil, status.Error(codes.Internal, "internal error")
    }
	return &manager1.UpdatePriceResponse{
		IsUpdate: true,
	}, nil
}

func (s *serverAPI) Get(ctx context.Context, req *manager1.GetPriceRequest) (*manager1.GetPriceResponse, error) {
	// Handle request
	err := m.HandleGet(req)
    if err!= nil{
        return nil, status.Error(codes.InvalidArgument, err.Error())
    }
    // Get a record.
    rec, err := s.p.Get(ctx, models.GetPrice{
        ID: req.UserId,
    })
    if err!= nil { 
		if errors.Is(err, price.ErrNotFound) {
            return nil, status.Error(codes.NotFound, "Entity not found")
        }
        return nil, status.Error(codes.Internal, "internal error")
    }
    return &manager1.GetPriceResponse{
        Price: rec.Price,
    }, nil
}