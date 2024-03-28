package price

import (
	"context"
	"errors"
	"log/slog"

	models "server/server/internal/models/price"
	"server/server/internal/storage"
)

// go:generate go run github.com/vektra/mockery/v2@v2.42.0 --name=PriceRepository
type PriceRepository interface {
	Create(
		ctx context.Context,
		rec *models.CreatePrice,
	) error
	Update(
		ctx context.Context,
		rec *models.UpdatePrice,
	) error
	Get(
		ctx context.Context,
		rec *models.GetPrice,
	) (price models.Prices, err error)
}

var (
	ErrInternal           = errors.New("internal error")
	ErrNotFound           = errors.New("entity is not found")
	ErrAlreadyExists      = errors.New("entity already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type Prices struct {
	log *slog.Logger
	rep PriceRepository
}

func New(rep PriceRepository, log *slog.Logger) *Prices {
	return &Prices{
		log: log,
		rep: rep,
	}
}

func (p *Prices) Create(
	ctx context.Context,
	rec models.CreatePrice,
) error {
	op := "Price. Create"
	p.log.Info(op)

	if err := p.rep.Create(ctx, &rec); err != nil {
		if errors.Is(err, storage.ErrAlreadyExists) {
			p.log.Error("%s: %v", op, err)
			return ErrAlreadyExists
		}
		p.log.Error("%s: %v", op, err)
		return err
	}
	return nil
}

func (p *Prices) Update(
	ctx context.Context,
	rec models.UpdatePrice,
) error {
	op := "Price. Update"
	p.log.Info(op)

	if err := p.rep.Update(ctx, &rec); err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			p.log.Error("%s: %v", op, err)
			return ErrNotFound
		}
		p.log.Error("%s: %v", op, err)
		return err
	}
	return nil
}

func (p *Prices) Get(
	ctx context.Context,
	rec models.GetPrice,
) (models.Prices, error) {
	op := "Price. Get"
	p.log.Info(op)

	price, err := p.rep.Get(ctx, &rec)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			p.log.Error("%s: %v", op, err)
			return models.Prices{}, ErrNotFound
		}
		p.log.Error("%s: %v", op, err)
		return models.Prices{}, err
	}
	return price, nil
}
