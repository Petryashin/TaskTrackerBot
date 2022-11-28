package inline_action

import (
	task_usecase "github.com/petryashin/TaskTrackerBot/internal/usecase/task"
	user_usecase "github.com/petryashin/TaskTrackerBot/internal/usecase/user"
)

const (
	AddTaskActionText    = "Напишите текст задачи"
	RemoveTaskActionText = "Напишите номер задачи, которую нужно удалить"
)

type TaskInterface task_usecase.TaskInterface

type UserInterface user_usecase.UserInterface

type RedisCacheInterface interface {
	Set(key string, json string) error
	Get(key string) (string, error)
}
