package payments

import (
	"context"
	"log/slog"

	models "server/server/internal/models/work"
)


type RepositoryWork interface {
	Create(
		ctx context.Context,
        rec models.CreateWork,
	)(bool, error)
	Update(
		ctx context.Context,
        rec models.UpdateWork,
	)(bool, error)
	GetAll(
		ctx context.Context,
        rec []*models.Work,
	)error
	GetAllByEmail(
		ctx context.Context,
        rec []*models.Work,
	)error
	Delete(
		ctx context.Context,
        id int64,
	)(bool, error)
} 

type Works struct {
	log *slog.Logger
	rep RepositoryWork
}

func New(rep RepositoryWork, log *slog.Logger) *Works {
	return &Works{
        log: log,
        rep: rep,
    }
}

func (w *Works) Create(
	ctx context.Context,
	rec models.CreateWork,
)(bool, error){
	return true, nil
}
func (w *Works) Update(
	ctx context.Context,
	rec models.UpdateWork,
)(bool, error){
	return true, nil
}
func (w *Works) Get(
	ctx context.Context,
	rec models.GetAllWork,
)([]models.Work, error){
	return []models.Work{}, nil
}
func (w *Works) GetByDate(
	ctx context.Context,
	rec models.GetAllWorkByDate,
)([]models.Work, error){
	return []models.Work{}, nil
}
func (w *Works) Delete(
	ctx context.Context,
	rec models.DeleteWork,
)(bool, error){
	return true, nil
}
