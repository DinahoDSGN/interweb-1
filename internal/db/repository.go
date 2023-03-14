package db

import (
	"context"
	"interweb/telegram-bot-service/pkg/domain"
	"time"
)

type Repository interface {
	InsertUserRequest(ctx context.Context, req domain.UserRequest) (int64, error)
	GetDateFirstRequest(ctx context.Context, id int64) (time.Time, error)
	AggregateTotalRequests(ctx context.Context, id int64) ([]domain.TotalUserRequests, error)
	ListRequests(ctx context.Context, id int64, userRequests chan<- domain.UserRequest) error
}

type PostgresRepository struct {
	Repository
}

func NewPostgresRepository(repository Repository) *PostgresRepository {
	return &PostgresRepository{Repository: repository}
}
