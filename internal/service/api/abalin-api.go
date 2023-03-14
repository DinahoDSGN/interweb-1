package api

import (
	"context"
	"io"
	"log"
	"net/http"
)

type AbalinApi struct {
}

func NewAbalinApi() *AbalinApi {
	return &AbalinApi{}
}

func (a *AbalinApi) Get(ctx context.Context) ([]byte, error) {
	res, err := http.Get("https://nameday.abalin.net/api/V1/today")
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	log.Println("AKSDJKAOSJD")

	return body, nil
}
