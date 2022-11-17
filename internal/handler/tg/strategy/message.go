package tgstrategy

import (
	"fmt"
	"strconv"

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
	users      userInterface
	redisCache redisCacheInterface
}

func NewMessageStrategy(tasks taskInterface, users userInterface, redisCache redisCacheInterface) MessageStrategy {
	return MessageStrategy{tasks: tasks, users: users, redisCache: redisCache}
}

func (i MessageStrategy) Handle(dto tgdto.DTO) (tgbotapi.MessageConfig, error) {
	action, err := i.redisCache.Get(int64toA(dto.System.ChatId))
	if err != nil {
		action = list
	}
	switch action {
	case addTask:
		newMessageText := dto.System.MessageText
		err := i.tasks.Add(dto.User.Id, newMessageText)
		if err != nil {
			return tgbotapi.MessageConfig{}, err
		}

		i.setDefaultState(dto.System)

		return i.messageBuilder(dto)
	case rmTask:
		taskNumber, err := strconv.Atoi(dto.System.MessageText)

		if err != nil {
			return tgbotapi.NewMessage(dto.User.TgId, "Введите номер задачи"), err
		}
		//err = i.tasks.Remove(dto.User.TgId, taskNumber)
		tasksList, err := i.tasks.List(dto.User.Id)
		if taskNumber < 1 || taskNumber > len(tasksList) || err != nil {
			return tgbotapi.NewMessage(dto.User.TgId, "Введите корректный номер задачи"), err
		}
		taskForRemove := tasksList[taskNumber-1]
		err = i.tasks.Remove(taskForRemove.Id)
		if err != nil {
			return tgbotapi.NewMessage(dto.User.TgId, "Не удалось удалить задачу"), err
		}
		i.setDefaultState(dto.System)
		return i.messageBuilder(dto)
	case list:
		return i.messageBuilder(dto)
	default:
		return i.messageBuilder(dto)
	}
}

func (i MessageStrategy) messageBuilder(dto tgdto.DTO) (tgbotapi.MessageConfig, error) {
	messageText := "Мои задачи:\n"

	tasksList, err := i.tasks.List(dto.User.Id)
	if err != nil {
		return tgbotapi.MessageConfig{}, err
	}
	for i, task := range tasksList {
		messageText += fmt.Sprintf("%d.  %s \n", i+1, task.Text)
	}
	msg := tgbotapi.NewMessage(dto.User.TgId, messageText)

	msg.ReplyMarkup = numericKeyboard
	return msg, nil
}

func (i MessageStrategy) setDefaultState(dto tgdto.SystemDTO) error {
	return i.redisCache.Set(int64toA(dto.ChatId), list)
}
