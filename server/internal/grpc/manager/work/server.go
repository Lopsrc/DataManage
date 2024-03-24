package work

import (
	"context"

	manager1 "server/protos/gen/go/manager"
	models "server/server/internal/models/work"
	m "server/server/internal/middleware/manager/work"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


type WorkService interface{
	Create(
		ctx context.Context,
        rec models.CreateWork,
	)(bool, error)
	Update(
		ctx context.Context,
        rec models.UpdateWork,
	)(bool, error)
	Get(
		ctx context.Context,
        rec models.GetAllWork,
	)([]models.Work, error)
	GetByDate(
		ctx context.Context,
        rec models.GetAllWorkByDate,
	)([]models.Work, error)
	Delete(
		ctx context.Context,
        rec models.DeleteWork,
	)(bool, error)
}

type serverAPI struct {
	manager1.UnimplementedManagerWorkServer
	w WorkService
}

func Register(gRPC *grpc.Server, work WorkService) {
	manager1.RegisterManagerWorkServer(gRPC, &serverAPI{w: work})
}

func (s *serverAPI) Create(ctx context.Context, req *manager1.CreateWorkRequest) (*manager1.CreateWorkResponse, error){
	// Handle requests.
	strErr, err:= m.HandleCreate(req)
	if err != nil{
		return nil, status.Error(codes.InvalidArgument, strErr)
	}
	// Create a new record
	isCreate, err := s.w.Create(ctx, models.CreateWork{
		Name: req.Name,
		Date: req.Date,
        Price: req.Price,
        Time: req.Time,
        Penalty: req.Penalty,
		UserID: req.UserId,
	})
	if err != nil { // FIXME:Errors
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &manager1.CreateWorkResponse{
		IsCreate: isCreate,
	}, nil
}

func (s *serverAPI) Update(ctx context.Context, req *manager1.UpdateWorkRequest) (*manager1.UpdateWorkResponse, error){
	// Handle requests.
	strErr, err:= m.HandleUpdate(req)
	if err != nil{
		return nil, status.Error(codes.InvalidArgument, strErr)
	}
	// Update an existing record.
	isUpdate, err := s.w.Update(ctx, models.UpdateWork{
		ID: req.Id,
		Name: req.Name,
		Date: req.Date,
        Price: req.Price,
        Time: req.Time,
        Penalty: req.Penalty,
	})
	if err != nil { // FIXME:Errors
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &manager1.UpdateWorkResponse{
		IsUpdate: isUpdate,
	}, nil
}

func (s *serverAPI) GetAll(ctx context.Context, req *manager1.GetWorkRequest) (*manager1.GetAllWorkResponse, error) {
	// Handle requests.
	strErr, err:= m.HandleGet(req)
	if err != nil{
		return nil, status.Error(codes.InvalidArgument, strErr)
	}
	// Get all records.
	records, err := s.w.Get(ctx, models.GetAllWork{
		UserID: req.UserId,
		Name: req.Name,
	})
	if err != nil { // FIXME:Errors
		return nil, status.Error(codes.Internal, "internal error")
	}
	sliceWorks := []*manager1.GetWorkResponse{}
	for _, r := range records {
		sliceWorks = append(sliceWorks, &manager1.GetWorkResponse{
            Id: r.ID,
            Name: r.Name,
            Date: r.Date,
            Price: r.Price,
            Time: r.Time,
            Penalty: r.Penalty,
			UserId: r.UserID,
        })
    }
	return &manager1.GetAllWorkResponse{
		ListWorks: sliceWorks,
	}, nil
}

func (s *serverAPI) GetAllByDate(ctx context.Context, req *manager1.GetByDateWorkRequest) (*manager1.GetAllWorkResponse, error) {
	// Handle requests.
	strErr, err:= m.HandleGetByDate(req)
	if err != nil{
		return nil, status.Error(codes.InvalidArgument, strErr)
	}
	// Get all records.
	records, err := s.w.GetByDate(ctx, models.GetAllWorkByDate{
		UserID: req.UserId,
		Name: req.Name,
		Date: req.Date,
	})
	if err != nil { // FIXME:Errors
		return nil, status.Error(codes.Internal, "internal error")
	}
	sliceWorks := []*manager1.GetWorkResponse{}
	for _, r := range records {
		sliceWorks = append(sliceWorks, &manager1.GetWorkResponse{
            Id: r.ID,
            Name: r.Name,
            Date: r.Date,
            Price: r.Price,
            Time: r.Time,
            Penalty: r.Penalty,
			UserId: r.UserID,
        })
    }
	return &manager1.GetAllWorkResponse{
		ListWorks: sliceWorks,
	}, nil
}

func (s *serverAPI) Delete(ctx context.Context, req *manager1.DeleteWorkRequest) (*manager1.DeleteWorkResponse, error) {
	// Handle requests.
	strErr, err := m.HandleDelete(req)
	if err != nil{
		return nil, status.Error(codes.InvalidArgument, strErr)
	}
	// Delete a record.
	isDel, err := s.w.Delete(ctx, models.DeleteWork{
		ID: req.Id,
	})
	if err != nil { // FIXME:Errors
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &manager1.DeleteWorkResponse{
		IsDel: isDel,
	}, nil
}