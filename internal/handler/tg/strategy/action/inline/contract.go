package inline_action

import (
	"github.com/petryashin/TaskTrackerBot/internal/domain/entity/task"
	"github.com/petryashin/TaskTrackerBot/internal/domain/entity/user"
)

const (
	AddTaskActionText    = "Напишите текст задачи"
	RemoveTaskActionText = "Напишите номер задачи, которую нужно удалить"
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
