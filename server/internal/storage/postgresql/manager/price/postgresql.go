package postgresqlprice

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	models "server/server/internal/models/price"
	"server/server/internal/storage"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct{
	log *slog.Logger
	client *pgxpool.Pool
}

func New(client *pgxpool.Pool, log *slog.Logger) *Repository {
	return &Repository{
		log: log,
		client: client,
	}
}

func (rep *Repository) Create(
	ctx context.Context,
	rec *models.CreatePrice,
)error{
	query := "INSERT INTO prices (user_id, price) VALUES ($1, $2)"

	_, err := rep.client.Exec(ctx, query, rec.ID, rec.Price)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			if pgErr.Code == storage.CodeAlreadyExists{
				return storage.ErrAlreadyExists
			}
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			return newErr
		}
		return err
	}

	return nil
}
func (rep *Repository) Update(
	ctx context.Context,
	rec *models.UpdatePrice,
)error{
	query := "UPDATE prices SET price = $1 WHERE user_id = $2"

	c, err := rep.client.Exec(ctx, query, rec.Price, rec.ID)
	if err!= nil {
		var pgErr *pgconn.PgError
        if errors.As(err, &pgErr) {
            pgErr = err.(*pgconn.PgError)
            newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			return newErr
        }
        return err
    }else if c.RowsAffected() == 0 {
		return storage.ErrNotFound
	}
	return nil
}
func (rep *Repository) Get(
	ctx context.Context,
	rec *models.GetPrice,
)(price models.Prices, err error){
	query := "SELECT user_id, price FROM prices WHERE user_id = $1"

	if err := rep.client.QueryRow(ctx, query, rec.ID).Scan(&price.ID, &price.Price); err!= nil {
		var pgErr *pgconn.PgError
        if errors.As(err, &pgErr) {
            pgErr = err.(*pgconn.PgError)
            newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			return models.Prices{}, newErr
        }else if err.Error() == storage.CodeNotFound{
			return models.Prices{}, storage.ErrNotFound
		}
		return models.Prices{}, err
	}

	return price, nil
}