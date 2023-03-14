package transport

import (
	"context"
	"interweb/telegram-bot-service/internal/transport/bot"
)

type Transport struct {
	bot bot.Bot
}

func NewTransport(bot bot.Bot) *Transport {
	return &Transport{bot: bot}
}

func (t *Transport) Listen(ctx context.Context) {
	go t.bot.ListenCommands(ctx)
}
