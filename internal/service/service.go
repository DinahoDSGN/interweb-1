package service

import (
	"context"
	"interweb/telegram-bot-service/internal/db"
	"interweb/telegram-bot-service/internal/service/api"
	"interweb/telegram-bot-service/pkg/domain"
	"time"
)

type Service struct {
	repo db.Repository
}

func NewService(repo db.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetDateFirstRequest(ctx context.Context, id int64) (time.Time, error) {
	datetime, err := s.repo.GetDateFirstRequest(ctx, id)
	if err != nil {
		return time.Time{}, err
	}

	return datetime, nil
}

func (s *Service) GetTotalRequests(ctx context.Context, id int64) ([]domain.TotalUserRequests, error) {
	totals, err := s.repo.AggregateTotalRequests(ctx, id)
	if err != nil {
		return []domain.TotalUserRequests{}, err
	}

	return totals, nil
}

func (s *Service) ListRequests(ctx context.Context, id int64, userRequests chan<- domain.UserRequest) error {
	err := s.repo.ListRequests(ctx, id, userRequests)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetDataByCommand(ctx context.Context, command string) ([]byte, error) {
	apis, err := api.GetApi(command)
	if err != nil {
		return nil, err
	}

	data, err := apis.Get(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Service) InsertRequest(ctx context.Context, req domain.UserRequest) (int64, error) {
	id, err := s.repo.InsertUserRequest(ctx, req)
	if err != nil {
		return 0, err
	}

	return id, nil
}
