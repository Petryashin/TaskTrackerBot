package tgstrategy

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
)

type InlineStrategy struct {
	tasks      taskInterface
	users      userInterface
	redisCache redisCacheInterface
}

func NewInlineStrategy(tasks taskInterface, users userInterface, redisCache redisCacheInterface) InlineStrategy {
	return InlineStrategy{tasks: tasks, users: users, redisCache: redisCache}
}

func (i InlineStrategy) Handle(dto tgdto.DTO) (tgbotapi.MessageConfig, error) {
	state := dto.System.MessageText
	i.setState(dto, state)
	switch state {
	case addTask:
		msg := tgbotapi.NewMessage(dto.System.ChatId, "Напишите текст задачи")
		return msg, nil
	case rmTask:
		msg := tgbotapi.NewMessage(dto.System.ChatId, "Напишите номер задачи, которую нужно удалить")
		return msg, nil
	}
	msg := tgbotapi.NewMessage(dto.System.ChatId, dto.System.MessageText)

	return msg, nil
}

func (i InlineStrategy) setDefaultState(dto tgdto.DTO) error {
	return i.redisCache.Set(int64toA(dto.System.ChatId), list)
}

func (i InlineStrategy) setState(dto tgdto.DTO, state string) error {
	return i.redisCache.Set(int64toA(dto.System.ChatId), state)
}
