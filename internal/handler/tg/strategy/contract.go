package tgstrategy

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/petryashin/TaskTrackerBot/internal/domain/entity/task"
	"github.com/petryashin/TaskTrackerBot/internal/domain/entity/user"
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
)

const addTask, rmTask, list = "addTask", "rmTask", "list"

type strategy interface {
	Handle(dto tgdto.DTO) (tgbotapi.MessageConfig, error)
}

type Strategies []strategy

type taskInterface interface {
	Add(userId int64, message string) (err error)
	Remove(taskId int64) error
	List(chatId int64) ([]task.Task, error)
}

type userInterface interface {
	Create(chatId int64, name string) (user.User, error)
	Remove(userId int64) error
	List() ([]user.User, error)
	FindOne(id int64) (user.User, error)
	FindOneByTgId(id int64) (user.User, error)
}

type redisCacheInterface interface {
	Set(key string, json string) error
	Get(key string) (string, error)
}
