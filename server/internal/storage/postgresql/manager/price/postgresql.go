package postgresqlprice

import (
	"context"

	models "server/server/internal/models/price"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct{
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

func (rep *Repository) Create(
	ctx context.Context,
	rec models.CreatePrice,
)(bool, error){
	return true, nil
}
func (rep *Repository) Update(
	ctx context.Context,
	rec models.UpdatePrice,
)(bool, error){
	return true, nil
}
func (rep *Repository) Get(
	ctx context.Context,
	rec *models.Prices,
)(bool, error){
	return true, nil
}
