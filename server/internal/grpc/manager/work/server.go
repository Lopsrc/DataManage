package work

import (
	"context"
	"errors"
	"fmt"
	"time"

	manager1 "server/protos/gen/go/manager"
	m "server/server/internal/middleware/manager/work"
	models "server/server/internal/models/work"
	work "server/server/internal/services/work"

	"github.com/jackc/pgtype"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
// go:generate go run github.com/vektra/mockery/v2@v2.42.0 --name=WorkService
type WorkService interface {
	Create(
		ctx context.Context,
		rec models.CreateWork,
	) error
	Update(
		ctx context.Context,
		rec models.UpdateWork,
	) error
	Get(
		ctx context.Context,
		rec models.GetAllWork,
	) ([]models.Work, error)
	GetByDate(
		ctx context.Context,
		rec models.GetAllWorkByDate,
	) ([]models.Work, error)
	Delete(
		ctx context.Context,
		rec models.DeleteWork,
	) error
}

type serverAPI struct {
	manager1.UnimplementedManagerWorkServer
	w WorkService
}
// Register registers the gRPC server for the work service.
func Register(gRPC *grpc.Server, work WorkService) {
	manager1.RegisterManagerWorkServer(gRPC, &serverAPI{w: work})
}
// Create creates a new work record.
// If the record already exists, ErrAlreadyExists is returned.
// If the request is invalid, ErrInvalidRequest is returned.
// If an internal error occurs, ErrInternal is returned.
func (s *serverAPI) Create(ctx context.Context, req *manager1.CreateWorkRequest) (*manager1.CreateWorkResponse, error) {
	// Handle requests.
	if err := m.HandleCreate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// Preparing the date.
	date, err := PrepareDate(req.GetDate())
	if err!= nil {
        return nil, status.Error(codes.InvalidArgument, "invalid date")
    }
	// Create a new record
	err = s.w.Create(ctx, models.CreateWork{
		Name:    req.Name,
		Date:    date,
		Time:    req.Time,
		Penalty: req.Penalty,
		UserID:  req.UserId,
	})
	if err != nil { 
		if errors.Is(err, work.ErrAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "already exists")
		}else if errors.Is(err, work.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "entity not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &manager1.CreateWorkResponse{
		IsCreate: true,
	}, nil
}
// Update updates an existing work record.
// If the record does not exist, ErrNotFound is returned.
// If the request is invalid, ErrInvalidRequest is returned.
// If an internal error occurs, ErrInternal is returned.
func (s *serverAPI) Update(ctx context.Context, req *manager1.UpdateWorkRequest) (*manager1.UpdateWorkResponse, error) {
	// Handle requests.
	if err := m.HandleUpdate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// Preparing the date.
	date, err := PrepareDate(req.GetDate())
	if err!= nil {
        return nil, status.Error(codes.InvalidArgument, "invalid date")
    }
	// Update an existing record.
	err = s.w.Update(ctx, models.UpdateWork{
		ID:      req.Id,
		Name:    req.Name,
		Date:    date,
		Time:    req.Time,
		Penalty: req.Penalty,
	})
	if err != nil { 
		if errors.Is(err, work.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "entity not found")
		}  
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &manager1.UpdateWorkResponse{
		IsUpdate: true,
	}, nil
}
// GetAll returns all work records that match the given filter criteria.
// If no criteria are specified, all records are returned.
// If an error occurs, an internal server error is returned.
func (s *serverAPI) GetAll(ctx context.Context, req *manager1.GetWorkRequest) (*manager1.GetAllWorkResponse, error) {
	// Handle requests.
	if err := m.HandleGet(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// Get all records.
	records, err := s.w.Get(ctx, models.GetAllWork{
		UserID: req.UserId,
		Name:   req.Name,
	})
	if err != nil { 
		if errors.Is(err, work.ErrNotFound) {
            return nil, status.Error(codes.NotFound, "entity not found")
        }
		return nil, status.Error(codes.Internal, "internal error")
	}
	sliceWorks := []*manager1.GetWorkResponse{}
	for _, r := range records {
		sliceWorks = append(sliceWorks, &manager1.GetWorkResponse{
			Id:      r.ID,
			Name:    r.Name,
			Date:    GetDate(r.Date.Time),
			Price:   r.Price,
			Time:    r.Time,
			Penalty: r.Penalty,
			UserId:  r.UserID,
		})
	}
	return &manager1.GetAllWorkResponse{
		ListWorks: sliceWorks,
	}, nil
}
// GetAllByDate returns all work records that match the given filter criteria.
// If no criteria are specified, all records are returned.
// If an error occurs, an internal server error is returned.
func (s *serverAPI) GetAllByDate(ctx context.Context, req *manager1.GetByDateWorkRequest) (*manager1.GetAllWorkResponse, error) {
	// Handle requests.
	if err := m.HandleGetByDate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
    // Get all records.
    records, err := s.w.GetByDate(ctx, models.GetAllWorkByDate{
        UserID: req.UserId,
        Name:   req.Name,
        Date:   req.Date,
    })
    if err != nil { 
        if errors.Is(err, work.ErrNotFound) {
            return nil, status.Error(codes.NotFound, "entity not found")
        }
        return nil, status.Error(codes.Internal, "internal error")
    }
	
	sliceWorks := []*manager1.GetWorkResponse{}
	for _, r := range records {
		sliceWorks = append(sliceWorks, &manager1.GetWorkResponse{
			Id:      r.ID,
			Name:    r.Name,
			Date:    GetDate(r.Date.Time),
			Price:   r.Price,
			Time:    r.Time,
			Penalty: r.Penalty,
			UserId:  r.UserID,
		})
	}
	return &manager1.GetAllWorkResponse{
		ListWorks: sliceWorks,
	}, nil
}
// Delete deletes a work record.
// If the record does not exist, ErrNotFound is returned.
// If the request is invalid, ErrInvalidRequest is returned.
// If an internal error occurs, ErrInternal is returned.
func (s *serverAPI) Delete(ctx context.Context, req *manager1.DeleteWorkRequest) (*manager1.DeleteWorkResponse, error) {
	// Handle requests.
	if err := m.HandleDelete(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// Delete a record.
	err := s.w.Delete(ctx, models.DeleteWork{
		ID: req.Id,
	})
	if err != nil { 
		if errors.Is(err, work.ErrNotFound) {
            return nil, status.Error(codes.NotFound, "entity not found")
        }
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &manager1.DeleteWorkResponse{
		IsDel: true,
	}, nil
}
// PrepareDate takes a date string in the format "2006-01-02" and returns a pgtype.Date struct.
func PrepareDate(date string) (pgtype.Date, error){
	format := "2006-01-02" 
	t, err := time.Parse(format, date)
	if err!= nil {
        return pgtype.Date{}, err
    }	
	return pgtype.Date{
		Time: t,
		Status: 0,
        InfinityModifier: 0,
	}, nil
}
// GetDate returns the date in the format "January 2, 2006"
func GetDate(time time.Time) string {
	return fmt.Sprintf("%s %d, %d", time.Month().String(), time.Day(), time.Year())
}