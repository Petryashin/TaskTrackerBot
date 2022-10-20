package tgstrategy

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Добавить задачу", addTask),
		tgbotapi.NewInlineKeyboardButtonData("Удалить задачу", rmTask),
	),
)

type MessageStrategy struct {
	tasks      taskInterface
	redisCache redisCacheInterface
}

func NewMessageStrategy(tasks taskInterface, redisCache redisCacheInterface) MessageStrategy {
	return MessageStrategy{tasks: tasks, redisCache: redisCache}
}

func (i MessageStrategy) Handle(dto tgdto.Dto) (tgbotapi.MessageConfig, error) {
	action, err := i.redisCache.Get(int64toA(dto.ChatId))
	if err != nil {
		action = list
	}
	switch action {
	case addTask:
		newMessageText := dto.MessageText
		err := i.tasks.Add(newMessageText)
		if err != nil {
			return tgbotapi.MessageConfig{}, err
		}

		i.setDefaultState(dto)

		return i.messageBuilder(dto)
	case rmTask:
		i.setDefaultState(dto)

		return i.messageBuilder(dto)
	case list:
		return i.messageBuilder(dto)
	default:
		return i.messageBuilder(dto)
	}
}

func (i MessageStrategy) messageBuilder(dto tgdto.Dto) (tgbotapi.MessageConfig, error) {
	messageText := "Мои задачи:\n"
	tasksList, err := i.tasks.List()
	if err != nil {
		return tgbotapi.MessageConfig{}, err
	}
	for i, task := range tasksList {
		messageText += fmt.Sprintf("%d.  %s \n", i+1, task.Text)
	}
	msg := tgbotapi.NewMessage(dto.ChatId, messageText)

	msg.ReplyMarkup = numericKeyboard
	return msg, nil
}

func (i MessageStrategy) setDefaultState(dto tgdto.Dto) error {
	return i.redisCache.Set(int64toA(dto.ChatId), list)
}
