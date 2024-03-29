package work

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
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

func Register(gRPC *grpc.Server, work WorkService) {
	manager1.RegisterManagerWorkServer(gRPC, &serverAPI{w: work})
}

func (s *serverAPI) Create(ctx context.Context, req *manager1.CreateWorkRequest) (*manager1.CreateWorkResponse, error) {
	// Handle requests.
	strErr, err := m.HandleCreate(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, strErr)
	}
	// Preparing the date.
	date, err := PrepareDate(req.GetDate())
	if err!= nil {
        return nil, status.Error(codes.InvalidArgument, "invalid date")
    }
	// Create a new record
	err = s.w.Create(ctx, models.CreateWork{
		Name:    req.Name,
		Date:    pgtype.Date{
			Time: 				date,
			Status: 			0,
			InfinityModifier: 	0,
		},
		Time:    req.Time,
		Penalty: req.Penalty,
		UserID:  req.UserId,
	})
	if err != nil { 
		if errors.Is(err, work.ErrAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "already exists")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &manager1.CreateWorkResponse{
		IsCreate: true,
	}, nil
}

func (s *serverAPI) Update(ctx context.Context, req *manager1.UpdateWorkRequest) (*manager1.UpdateWorkResponse, error) {
	// Handle requests.
	strErr, err := m.HandleUpdate(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, strErr)
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
		Date:    pgtype.Date{
			Time: 				date,
			Status: 			0,
			InfinityModifier: 	0,
		},
		Time:    req.Time,
		Penalty: req.Penalty,
	})
	if err != nil { 
		if errors.Is(err, work.ErrNotFound) {
            return nil, status.Error(codes.NotFound, "not found")
        }
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &manager1.UpdateWorkResponse{
		IsUpdate: true,
	}, nil
}

func (s *serverAPI) GetAll(ctx context.Context, req *manager1.GetWorkRequest) (*manager1.GetAllWorkResponse, error) {
	// Handle requests.
	strErr, err := m.HandleGet(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, strErr)
	}
	// Get all records.
	records, err := s.w.Get(ctx, models.GetAllWork{
		UserID: req.UserId,
		Name:   req.Name,
	})
	if err != nil { 
		if errors.Is(err, work.ErrNotFound) {
            return nil, status.Error(codes.NotFound, "not found")
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

func (s *serverAPI) GetAllByDate(ctx context.Context, req *manager1.GetByDateWorkRequest) (*manager1.GetAllWorkResponse, error) {
	// Handle requests.
	strErr, err := m.HandleGetByDate(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, strErr)
	}
    // Get all records.
    records, err := s.w.GetByDate(ctx, models.GetAllWorkByDate{
        UserID: req.UserId,
        Name:   req.Name,
        Date:   req.Date,
    })
    if err != nil { 
        if errors.Is(err, work.ErrNotFound) {
            return nil, status.Error(codes.NotFound, "not found")
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

func (s *serverAPI) Delete(ctx context.Context, req *manager1.DeleteWorkRequest) (*manager1.DeleteWorkResponse, error) {
	// Handle requests.
	strErr, err := m.HandleDelete(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, strErr)
	}
	// Delete a record.
	err = s.w.Delete(ctx, models.DeleteWork{
		ID: req.Id,
	})
	if err != nil { 
		if errors.Is(err, work.ErrNotFound) {
            return nil, status.Error(codes.NotFound, "not found")
        }
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &manager1.DeleteWorkResponse{
		IsDel: true,
	}, nil
}

func PrepareDate(date string) (time time.Time, err error){
	arrDate := strings.Split(date, "-")
	year, err := strconv.Atoi(arrDate[0])
	if err!= nil {
        return time, err
    }
	month, err := strconv.Atoi(arrDate[1])
	if err!= nil {
        return time, err
    }
	day, err := strconv.Atoi(arrDate[2])
	if err!= nil {
        return time, err
    }
	time = time.AddDate(year-1, month-1, day-1)
	return time, nil
}

func GetDate(time time.Time) string {
	
	return fmt.Sprintf("%s %d, %d", time.Month().String(), time.Day(), time.Year())
}