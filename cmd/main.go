package main

import (
	"context"
	"fmt"
	"interweb/telegram-bot-service/internal/config"
	"interweb/telegram-bot-service/internal/db"
	"interweb/telegram-bot-service/internal/db/postgres"
	"interweb/telegram-bot-service/internal/service"
	"interweb/telegram-bot-service/internal/transport"
	"interweb/telegram-bot-service/internal/transport/bot"
	"interweb/telegram-bot-service/pkg/database"
	"interweb/telegram-bot-service/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

// init logger
func init() {
	logger.SetLogger(logger.NewZapLogger())
}

func main() {
	// init configs
	var cfg, err = config.New()
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to initialize config, err: %v", err))
	}

	pg, err := database.NewPostgres(cfg)
	if err != nil {
		logger.Error(err.Error())
	}
	defer func(pg *database.Postgres) {
		if err = pg.Close(); err != nil {
			logger.Fatal(fmt.Sprintf("failed to close connection, err: %v", err.Error()))
		}
	}(pg)

	// init dep-s
	var pgRepo = postgres.NewPostgres(pg)
	var r = db.NewPostgresRepository(pgRepo)
	var s = service.NewService(r)
	b, err := bot.NewTelegramBot(s, cfg)
	if err != nil {
		logger.Fatal(err.Error())
	}
	var t = transport.NewTransport(b)
	go t.Listen(context.Background())

	// graceful shutdown
	quit := make(chan os.Signal, 1) // TODO <- доделать
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	osSignal := <-quit

	logger.Info(fmt.Sprintf("program shutdown... call_type: %v", osSignal))
}
