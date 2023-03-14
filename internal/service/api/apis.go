package api

import (
	"context"
	"interweb/telegram-bot-service/pkg/domain"
)

type Api interface {
	Get(ctx context.Context) ([]byte, error)
}

func GetApi(service string) (Api, error) {
	switch service {
	case domain.CommandCryptocurrency:
		return NewCryptocurrencyApi(), nil
	case domain.CommandAbalin:
		return NewAbalinApi(), nil
	case domain.CommandWeather:
		return NewWeatherApi(), nil
	default:
		return nil, domain.ErrNoSuchApiServiceCommand
	}
}
