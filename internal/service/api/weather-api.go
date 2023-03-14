package api

import (
	"context"
	"io"
	"log"
	"net/http"
)

type WeatherApi struct {
}

func NewWeatherApi() *WeatherApi {
	return &WeatherApi{}
}

func (a *WeatherApi) Get(ctx context.Context) ([]byte, error) {
	res, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=43.26&longitude=76.93&hourly=rain")
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	log.Println("AKSDJKAOSJD")

	return body, nil
}
