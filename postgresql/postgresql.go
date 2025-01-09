package postgresql

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"os"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func LoadUrl() string {
	pgUserName := os.Getenv("PG_NAME")
	pgPassword := os.Getenv("PG_PASSWORD")
	pgHost := os.Getenv("PG_HOST")
	pgPort := os.Getenv("PG_PORT")
	pgDatabase := os.Getenv("PG_DATABASE")

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", pgUserName, pgPassword, pgHost, pgPort, pgDatabase)
	return url
}

// NewPostgresClient NewClient "postgres://username:password@localhost:5432/database_name"
func NewPostgresClient() (*pgx.Conn, error) {
	url := LoadUrl()
	connect, errConnect := pgx.Connect(context.Background(), url)

	if errConnect != nil {
		return nil, errors.New(
			fmt.Sprintf("failed to connect database. due to error: %v, errConnect"))
	}
	return connect, nil
}
