package domain

import "errors"

var ErrTimeoutCases = []string{"TLS handshake timeout", "connect: connection refused, api response"}

type ErrorMessage struct {
	Text error
	Code int
}

var (
	ErrNoData                  = errors.New("no data")
	ErrNoSuchApiServiceCommand = errors.New("no such api service command")
)
