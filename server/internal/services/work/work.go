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

func New(rep WorkRepository, log *slog.Logger) *Works {
	return &Works{
		log: log,
		rep: rep,
	}
}

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

func SortByMonth(w *[]models.Work, month string) (works []models.Work) {
	for _, i := range *w {
		if strings.EqualFold(i.Date.Time.Month().String(), month) {
			works = append(works, i)
		}
	}
	return
}
