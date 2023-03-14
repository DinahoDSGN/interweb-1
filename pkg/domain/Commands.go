package domain

import (
	"bytes"
)

const (
	InfoFirstRequest  = "info_first_request"
	InfoTotalRequests = "info_total_requests"
	InfoRequestList   = "info_request_list"
	Help              = "help"
)

const (
	CommandCryptocurrency = "cryptocurrency"
	CommandAbalin         = "abalin"
	CommandWeather        = "weather"
)

var CommandList CommandTypes = []string{CommandCryptocurrency, CommandAbalin, CommandWeather}

type CommandTypes []string

func (c CommandTypes) String() string {
	var o bytes.Buffer
	for _, s := range c {
		o.WriteString("/" + s + " ")
	}

	return o.String()
}

var InfoCommandList InfoCommandTypes = []string{InfoRequestList, InfoTotalRequests, InfoRequestList, Help}

type InfoCommandTypes []string

func (c InfoCommandTypes) String() string {
	var o bytes.Buffer
	for _, s := range c {
		o.WriteString("/" + s + " ")
	}

	return o.String()
}
