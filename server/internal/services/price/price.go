package price

import (
	"context"
	"log/slog"

	models "server/server/internal/models/price"

)

type RepositoryPrice interface {
	Create(
		ctx context.Context,
        rec models.CreatePrice,
	)(bool, error)
	Update(
		ctx context.Context,
        rec models.UpdatePrice,
	)(bool, error)
	Get(
		ctx context.Context,
        rec *models.Prices,
	)(bool, error)
}

type Prices struct {
	log *slog.Logger
	rep RepositoryPrice
}

func New(rep RepositoryPrice, log *slog.Logger) *Prices {
	return &Prices{
		log: log,
        rep: rep,
	}
}

func (p *Prices) Create(
	ctx context.Context,
	rec models.CreatePrice,
)(bool, error){
	return true, nil
}

func (p *Prices) Update(
	ctx context.Context,
	rec models.UpdatePrice,
)(bool, error){
	return true, nil
}

func (p *Prices) Get(
	ctx context.Context,
	rec models.GetPrice,
)(models.Prices, error){
	return models.Prices{}, nil
}