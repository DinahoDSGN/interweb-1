package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"interweb/telegram-bot-service/internal/config"
)

type Postgres struct {
	Conn *sql.DB
}

func NewPostgres(cfg config.Config) (*Postgres, error) {
	var url = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.PostgresUser, cfg.PostgresPass, cfg.PostgresHost, cfg.PostgresDB)
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Postgres{
		db,
	}, nil
}

func (r *Postgres) Close() error {
	if err := r.Conn.Close(); err != nil {
		return err
	}

	return nil
}
