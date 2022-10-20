package tgstrategy

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
)

type InlineStrategy struct {
	tasks      taskInterface
	redisCache redisCacheInterface
}

func NewInlineStrategy(tasks taskInterface, redisCache redisCacheInterface) InlineStrategy {
	return InlineStrategy{tasks: tasks, redisCache: redisCache}
}

func (i InlineStrategy) Handle(dto tgdto.Dto) (tgbotapi.MessageConfig, error) {
	state := dto.MessageText
	i.setState(dto, state)
	switch state {
	case addTask:
		msg := tgbotapi.NewMessage(dto.ChatId, "Напишите текст задачи")
		return msg, nil
	case rmTask:
		msg := tgbotapi.NewMessage(dto.ChatId, "Напишите номер задачи, которую нужно удалить")
		return msg, nil
	}
	msg := tgbotapi.NewMessage(dto.ChatId, dto.MessageText)

	return msg, nil
}

func (i InlineStrategy) setDefaultState(dto tgdto.Dto) error {
	return i.redisCache.Set(int64toA(dto.ChatId), list)
}

func (i InlineStrategy) setState(dto tgdto.Dto, state string) error {
	return i.redisCache.Set(int64toA(dto.ChatId), state)
}
