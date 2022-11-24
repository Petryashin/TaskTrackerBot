package message_action

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/petryashin/TaskTrackerBot/internal/domain/entity/task"
	"github.com/petryashin/TaskTrackerBot/internal/domain/entity/user"
	strategy_constant "github.com/petryashin/TaskTrackerBot/internal/handler/tg/strategy/constant"
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

type TaskInterface interface {
	Add(userId int64, message string) (err error)
	Remove(taskId int64) error
	List(chatId int64) ([]task.Task, error)
}

type UserInterface interface {
	Create(chatId int64, name string) (user.User, error)
	Remove(userId int64) error
	List() ([]user.User, error)
	FindOne(id int64) (user.User, error)
	FindOneByTgId(id int64) (user.User, error)
}

type RedisCacheInterface interface {
	Set(key string, json string) error
	Get(key string) (string, error)
}
