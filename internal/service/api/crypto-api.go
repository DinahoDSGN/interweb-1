package api

import "context"

type CryptocurrencyApi struct {
}

func NewCryptocurrencyApi() *CryptocurrencyApi {
	return &CryptocurrencyApi{}
}

func (a *CryptocurrencyApi) Get(ctx context.Context) ([]byte, error) {
	return []byte{1, 2, 3, 4, 5}, nil
}
