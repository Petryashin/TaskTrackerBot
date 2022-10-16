package task

import "github.com/petryashin/TaskTrackerBot/internal/storage/memory"

type cacheInterface interface {
	Add(message string) (err error)
	List() ([]memory.Task, error)
}

type Task struct {
	Text string
}
