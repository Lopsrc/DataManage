package postgresqlprice

import (
	"context"
	"errors"
	"fmt"

	models "server/server/internal/models/price"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct{
	
	client *pgxpool.Pool
}

func New(client *pgxpool.Pool) *Repository {
	return &Repository{
		client: client,
	}
}

func (rep *Repository) Create(
	ctx context.Context,
	rec models.CreatePrice,
)(bool, error){
	query := "INSERT INTO price (user_id, price) VALUES ($1, $2)"

	_, err := rep.client.Exec(ctx, query, rec.ID, rec.Price)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			return false, newErr
		}
		return false, err
	}

	return true, nil
}
func (rep *Repository) Update(
	ctx context.Context,
	rec models.UpdatePrice,
)(bool, error){
	query := "UPDATE price SET price = $1 WHERE user_id = $2"

	_, err := rep.client.Exec(ctx, query, rec.Price, rec.ID)
	if err!= nil {
		var pgErr *pgconn.PgError
        if errors.As(err, &pgErr) {
            pgErr = err.(*pgconn.PgError)
            newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
            return false, newErr
        }
        return false, err
    }
	return true, nil
}
func (rep *Repository) Get(
	ctx context.Context,
	rec *models.Prices,
)(bool, error){
	query := "SELECT price FROM price WHERE user_id = $1"

	if err := rep.client.QueryRow(ctx, query, rec.ID).Scan(&rec.Price); err!= nil {
		var pgErr *pgconn.PgError
        if errors.As(err, &pgErr) {
            pgErr = err.(*pgconn.PgError)
            newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
            return false, newErr
        }
		return false, err
	}

	return true, nil
}