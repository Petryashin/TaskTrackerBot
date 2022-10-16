package memory

type Storage struct {
	Cache *Cache
}

type Task struct {
	Text string
}

func New() (Storage, error) {
	tasks, err := mustParseCache()
	if err != nil {
		return Storage{}, err
	}

	return Storage{Cache: tasks}, nil
}

func (c Storage) Add(message string) (err error) {
	c.Cache.Tasks = append(c.Cache.Tasks, Task{message})
	c.Cache.mustPutCache()
	return
}

func (c Storage) List() ([]Task, error) {
	return c.Cache.Tasks, nil
}
