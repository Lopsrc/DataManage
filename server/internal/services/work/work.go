package workservice

import (
	"context"
	"errors"
	"log/slog"
	"strings"

	models "server/server/internal/models/work"
	"server/server/internal/storage"
)

// go:generate go run github.com/vektra/mockery/v2@v2.42.0 --name=WorkRepository
type WorkRepository interface {
	Create(
		ctx context.Context,
		rec *models.CreateWork,
	) error
	Update(
		ctx context.Context,
		rec *models.UpdateWork,
	) error
	GetAll(
		ctx context.Context,
		rec *models.GetAllWork,
	) (works []models.Work, err error)
	Delete(
		ctx context.Context,
		rec *models.DeleteWork,
	) error
}

var (
	ErrInternal           = errors.New("internal error")
	ErrNotFound           = errors.New("entity is not found")
	ErrAlreadyExists      = errors.New("entity already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type Works struct {
	log *slog.Logger
	rep WorkRepository
}
// New creates a new instance of Works.
func New(rep WorkRepository, log *slog.Logger) *Works {
	return &Works{
		log: log,
		rep: rep,
	}
}
// Create creates a new work record.
// If the work record already exists, ErrAlreadyExists is returned.
// If any other error occurs, it is returned.
func (w *Works) Create(
	ctx context.Context,
	rec models.CreateWork,
) error {
	op := "Work. Create"
	w.log.Info(op)

	if err := w.rep.Create(ctx, &rec); err != nil {
		if errors.Is(err, storage.ErrAlreadyExists) {
			w.log.Error("%s: %v", op, err)
			return ErrAlreadyExists
		} else if errors.Is(err, storage.ErrNotFound) {
			w.log.Error("%s: %v", op, err)
			return ErrNotFound
		}
		w.log.Error("%s: %v", op, err)
		return err
	}

	return nil
}
// Update updates an existing work record.
// If the work record does not exist, ErrNotFound is returned.
// If any other error occurs, it is returned.
func (w *Works) Update(
	ctx context.Context,
	rec models.UpdateWork,
) error {
	op := "Work. Update"
	w.log.Info(op)

	if err := w.rep.Update(ctx, &rec); err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			w.log.Error("%s: %v", op, err)
			return ErrNotFound
		}
		w.log.Error("%s: %v", op, err)
		return err
	}

	return nil
}
// Get retrieves all work records.
// If no records exist, an empty slice is returned.
// If an error occurs, it is returned.
func (w *Works) Get(
	ctx context.Context,
	rec models.GetAllWork,
) ([]models.Work, error) {
	op := "Work. Get"
	w.log.Info(op)

	works, err := w.rep.GetAll(ctx, &rec)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			w.log.Error("%s: %v", op, err)
			return []models.Work{}, ErrNotFound
		}
		w.log.Error("%s: %v", op, err)
		return []models.Work{}, err
	}
	return works, nil
}
// GetByDate retrieves all work records based on the date.
// If no records exist, an empty slice is returned.
// If an error occurs, it is returned.
func (w *Works) GetByDate(
	ctx context.Context,
	rec models.GetAllWorkByDate,
) ([]models.Work, error) {
	op := "Work. GetByDate"
	w.log.Info(op)

	works, err := w.rep.GetAll(ctx, &models.GetAllWork{
		UserID: rec.UserID,
		Name:   rec.Name,
	})
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			w.log.Error("%s: %v", op, err)
			return []models.Work{}, ErrNotFound
		}
		w.log.Error("%s: %v", op, err)
		return []models.Work{}, err
	}
	return SortByMonth(&works, rec.Date), nil
}
// Delete deletes an existing work record.
// If the work record does not exist, ErrNotFound is returned.
// If any other error occurs, it is returned.
func (w *Works) Delete(
	ctx context.Context,
	rec models.DeleteWork,
) error {
	op := "Work. Delete"
	w.log.Info(op)

	if err := w.rep.Delete(ctx, &rec); err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			w.log.Error("%s: %v", op, err)
			return ErrNotFound
		}
		w.log.Error("%s: %v", op, err)
		return err
	}

	return nil
}
// SortByMonth sorts a slice of Work structs by month.
// It returns a new slice of Work structs.
func SortByMonth(w *[]models.Work, month string) (works []models.Work) {
	for _, i := range *w {
		if strings.EqualFold(i.Date.Time.Month().String(), month) {
			works = append(works, i)
		}
	}
	return
}
