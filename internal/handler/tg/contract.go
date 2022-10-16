package tg

import "github.com/petryashin/TaskTrackerBot/internal/usecase/task"

type taskUsecase interface {
	Add(message string) error
	List() ([]task.Task, error)
}
