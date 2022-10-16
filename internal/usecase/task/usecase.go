package task

import "github.com/petryashin/TaskTrackerBot/internal/storage/memory"

type Usecase struct {
	cache cacheInterface
}

func New(cache cacheInterface) Usecase {
	return Usecase{cache: cache}
}

func (u Usecase) Add(message string) error {
	return u.cache.Add(message)
}

func (u Usecase) List() ([]Task, error) {
	memoryTasks, err := u.cache.List()
	if err != nil {
		return []Task{}, err
	}

	return convertTasks(memoryTasks), nil
}

func convertTasks(tasks []memory.Task) []Task {
	result := make([]Task, 0, len(tasks))
	for _, task := range tasks {
		result = append(result, Task{Text: task.Text})
	}
	return result
}
