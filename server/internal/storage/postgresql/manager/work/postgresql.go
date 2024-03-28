package postgresqlwork

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	models "server/server/internal/models/work"
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
	rec *models.CreateWork,
)error{
	query := "SELECT price FROM prices WHERE user_id = $1"

	if err := rep.client.QueryRow(ctx, query, rec.UserID).Scan(&rec.Price); err!= nil {
		var pgErr *pgconn.PgError
        if errors.As(err, &pgErr) {
            pgErr = err.(*pgconn.PgError)
            newErr := fmt.Errorf(fmt.Sprintf("Get. SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s" ,pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			return newErr
        }else if err.Error() == storage.CodeNotFound{
			return storage.ErrNotFound
		}
		return err
	}
	query = "INSERT INTO work(name, date, price, time, penalty, user_id) VALUES($1, $2, $3, $4, $5, $6)"

	_, err := rep.client.Exec(ctx, query, rec.Name, rec.Date.Time, rec.Price, rec.Time, rec.Penalty, rec.UserID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			if pgErr.Code == storage.CodeAlreadyExists{
				return storage.ErrAlreadyExists
			}
			newErr := fmt.Errorf(fmt.Sprintf("Create. SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			return newErr
		}
		return err
	}

	return nil
}
func (rep *Repository) Update(
	ctx context.Context,
	rec *models.UpdateWork,
)error{

	query := "UPDATE work SET name = $1, date = $2, price = $3, time = $4, penalty = $5 WHERE id = $6"

	c, err := rep.client.Exec(ctx, query, rec.Name, rec.Date.Time, rec.Price, rec.Time, rec.Penalty, rec.ID)
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
func (rep *Repository) GetAll(
	ctx context.Context,
	rec *models.GetAllWork,
)(works []models.Work, err error){

	query := "SELECT id, name, date, price, time, penalty FROM work WHERE name = $1 AND user_id = $2"

	r, err := rep.client.Query(ctx, query, rec.Name, rec.UserID)
	if err != nil{
		var pgErr *pgconn.PgError
        if errors.As(err, &pgErr) {
            pgErr = err.(*pgconn.PgError)
            newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			return nil, newErr
        }else if err.Error() == storage.CodeNotFound{
			return nil, storage.ErrNotFound
		}
        return nil, err
	}
	defer r.Close()

	for r.Next() {
		var work models.Work
        if err := r.Scan(&work.ID, &work.Name, &work.Date, &work.Price, &work.Time, &work.Penalty); err!= nil {
            return nil, err
        }
        works = append(works, work)
    }
	return 
}

func (rep *Repository) Delete(
	ctx context.Context,
	rec *models.DeleteWork,
)error{
	query := "DELETE FROM work WHERE id = $1"

	_, err := rep.client.Exec(ctx, query, rec.ID)
	if err!= nil {
		var pgErr *pgconn.PgError
        if errors.As(err, &pgErr) {
            pgErr = err.(*pgconn.PgError)
            newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			return newErr
        }else if err.Error() == storage.CodeNotFound{
			return  storage.ErrNotFound
		}
        return err
    }
	return nil
}