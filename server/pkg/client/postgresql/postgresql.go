package postgresql

import (
	"context"
	"fmt"
	"log"
	"server/server/internal/config"
	repeatable "server/server/pkg/utils"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, opts pgx.TxOptions) (pgx.Tx, error)
}
// NewClient creates a new PostgreSQL client.
func NewClient(ctx context.Context, maxAttempts int, cfg config.StorageConfig) (pool *pgxpool.Pool, err error) {
    // dsn is the data source name used to connect to the PostgreSQL database.
    dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	
    // DoWithTries executes the given function repeatedly until it returns no error or the maximum number of attempts is reached.
    err = repeatable.DoWithTries(func() error {
        // Create a context with a timeout of 5 seconds.
        ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
        defer cancel()

        // Connect to the PostgreSQL database using the given DSN.
        pool, err = pgxpool.Connect(ctx, dsn)
        if err!= nil {
            return err
        }

        return nil
    }, maxAttempts, 5*time.Second)

    // If an error occurs, log a fatal message and return.
    if err!= nil {
        log.Fatal("error do with tries postgresql")
    }

    return pool, nil
}
