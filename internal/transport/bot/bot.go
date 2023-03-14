package bot

import (
	"bytes"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"interweb/telegram-bot-service/internal/config"
	"interweb/telegram-bot-service/internal/service"
	"interweb/telegram-bot-service/pkg/domain"
	"interweb/telegram-bot-service/pkg/logger"
	"interweb/telegram-bot-service/tools"
	"strings"
)

type TelegramBot struct {
	service *service.Service
	bot     *tgbotapi.BotAPI
}

const minCut = 50

func NewTelegramBot(service *service.Service, cfg config.Config) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramBotToken)
	if err != nil {
		return nil, err
	}

	logger.Info(fmt.Sprintf("Authorized on account %s", bot.Self.UserName))

	return &TelegramBot{
		service: service,
		bot:     bot,
	}, nil
}

type Bot interface {
	ListenCommands(ctx context.Context)
}

func (b *TelegramBot) ListenCommands(ctx context.Context) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 100000

	for update := range b.bot.GetUpdatesChan(u) {
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		command := update.Message.Command()

		if strings.Contains(command, "info_") || command == domain.Help {
			go b.sendUserInfo(ctx, msg, command)
			continue
		}

		go b.sendDataFromApi(ctx, msg, command)
	}
}

func (b *TelegramBot) sendUserInfo(ctx context.Context, msg tgbotapi.MessageConfig, command string) {
	switch command {
	case domain.InfoFirstRequest:
		{
			datetime, err := b.service.GetDateFirstRequest(ctx, msg.ChatID)
			if err != nil {
				logger.Error(err.Error())
				return
			}

			msg.Text = datetime.Format("2006-01-02 15:04:05")
		}
	case domain.InfoTotalRequests:
		{
			totals, err := b.service.GetTotalRequests(ctx, msg.ChatID)
			if err != nil {
				logger.Error(err.Error())
				return
			}

			var totalsBuf bytes.Buffer
			for _, total := range totals {
				totalsBuf.WriteString(total.String() + "\n")
			}

			msg.Text = totalsBuf.String()
		}
	case domain.InfoRequestList:
		{
			var userRequests = make(chan domain.UserRequest, 10)
			go func() {
				err := b.service.ListRequests(ctx, msg.ChatID, userRequests)
				if err != nil {
					logger.Error(err.Error())
				}
			}()

			var userRequestsBuf bytes.Buffer
			for req := range userRequests {
				userRequestsBuf.WriteString(req.String() + "\n")
			}

			msg.Text = userRequestsBuf.String()
		}
	case domain.Help:
		msg.Text = domain.InfoCommandList.String()
	default:
		msg.Text = domain.InfoCommandList.String()
	}

	if _, err := b.bot.Send(msg); err != nil {
		logger.Error(err.Error())
		return
	}
}

func (b *TelegramBot) sendDataFromApi(ctx context.Context, msg tgbotapi.MessageConfig, command string) {
	res, err := b.service.GetDataByCommand(ctx, command)
	if err == domain.ErrNoSuchApiServiceCommand {
		msg.Text = domain.CommandList.String()

		if _, err := b.bot.Send(msg); err != nil {
			logger.Error(err.Error())
			return
		}

		return
	} else if err != nil {
		logger.Error(err)
		b.sendError(ctx, msg, domain.ErrorMessage{
			Text: err,
			Code: 0,
		})
		return
	}

	msg.Text = string(res)
	if len(msg.Text) > minCut {
		msg.Text = msg.Text[:minCut] + "..."
	}

	if res == nil {
		msg.Text = domain.ErrNoData.Error()
	}

	if _, err := b.bot.Send(msg); err != nil {
		logger.Error(err)
		return
	}

	id, err := b.service.InsertRequest(ctx, domain.UserRequest{
		ChatID:  msg.ChatID,
		Request: command,
		Result:  res,
	})
	if err != nil {
		logger.Error(err.Error())
		return
	}

	logger.Info(fmt.Sprintf("%d affected", id)) // TODO <- убрать потом
}

func (b *TelegramBot) sendError(ctx context.Context, msg tgbotapi.MessageConfig, errorMessage domain.ErrorMessage) {
	if tools.ContainsString(errorMessage.Text.Error(), domain.ErrTimeoutCases) {
		errorMessage.Code = 001
		errorMessage.Text = fmt.Errorf("it seems like service does not work. text: %s", errorMessage.Text)
	}

	msg.Text = fmt.Sprintf("An error occured: %v\napi response status code: %d", errorMessage.Text, errorMessage.Code)

	if _, err := b.bot.Send(msg); err != nil {
		logger.Error(err.Error())
		return
	}
}
