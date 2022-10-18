package tgstrategy

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
)

type InlineStrategy struct {
	cache cacheInterface
}

func NewInlineStrategy(cache cacheInterface) InlineStrategy {
	return InlineStrategy{cache: cache}
}

func (i InlineStrategy) Handle(dto tgdto.Dto) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(dto.ChatId, dto.MessageText)

	return msg
}
