package tgstrategy

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
	message_action "github.com/petryashin/TaskTrackerBot/internal/handler/tg/strategy/action/message"
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Добавить задачу", AddTask),
		tgbotapi.NewInlineKeyboardButtonData("Удалить задачу", RmTask),
	),
)

type MessageStrategy struct {
	actions    Actions
	tasks      TaskInterface
	users      UserInterface
	redisCache RedisCacheInterface
}

func NewMessageStrategy(tasks TaskInterface, users UserInterface, redisCache RedisCacheInterface) MessageStrategy {
	actions := createMessageStrategyActions(tasks, users, redisCache)
	return MessageStrategy{actions: actions, tasks: tasks, users: users, redisCache: redisCache}
}

func createMessageStrategyActions(tasks TaskInterface, users UserInterface, redisCache RedisCacheInterface) Actions {
	return Actions{
		Main:    message_action.NewDefaultAction(tasks, users, redisCache),
		AddTask: message_action.NewAddTaskAction(tasks, users, redisCache),
		RmTask:  message_action.NewRemoveTaskAction(tasks, users, redisCache),
	}
}

func (i MessageStrategy) Handle(dto tgdto.DTO) (tgbotapi.MessageConfig, error) {
	action, err := i.redisCache.Get(int64toA(dto.System.ChatId))
	if err != nil {
		action = Main
	}
	_, err = i.actions[action].Handle(dto)

	if err != nil {
		return tgbotapi.NewMessage(dto.User.TgId, err.Error()), nil
	} else {
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
	return i.redisCache.Set(int64toA(dto.ChatId), Main)
}
