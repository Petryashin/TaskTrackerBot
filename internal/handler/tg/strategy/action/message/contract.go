package message_action

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	strategy_constant "github.com/petryashin/TaskTrackerBot/internal/handler/tg/strategy/constant"
	task_usecase "github.com/petryashin/TaskTrackerBot/internal/usecase/task"
	user_usecase "github.com/petryashin/TaskTrackerBot/internal/usecase/user"
)

const (
	incorrectRemovingTaskNumberText = "Введите корректный номер задачи"
	MustBeNumeric                   = "Номер задачи должен быть числом"
	unsuccessfulRemovingTaskText    = "Не удалось удалить задачу"
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Добавить задачу", strategy_constant.AddTask),
		tgbotapi.NewInlineKeyboardButtonData("Удалить задачу", strategy_constant.RmTask),
	),
)

type TaskInterface task_usecase.TaskInterface

type UserInterface user_usecase.UserInterface

type RedisCacheInterface interface {
	Set(key string, json string) error
	Get(key string) (string, error)
}
