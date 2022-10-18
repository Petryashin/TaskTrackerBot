package tgstrategy

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
)

type MessageStrategy struct {
	cache cacheInterface
}

func NewMessageStrategy(cache cacheInterface) MessageStrategy {
	return MessageStrategy{cache: cache}
}

func (i MessageStrategy) Handle(dto tgdto.Dto) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(dto.ChatId, dto.MessageText)

	return msg
}
