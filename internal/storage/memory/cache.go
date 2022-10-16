package memory

type Storage struct {
	Cache *Cache
}

type Task struct {
	Text string
}

func New() Storage {
	return Storage{Cache: mustParseCache()}
}

func (c Storage) Add(message string) (err error) {
	c.Cache.Tasks = append(c.Cache.Tasks, Task{message})
	c.Cache.mustPutCache()
	return
}

func (c Storage) List() ([]Task, error) {
	return c.Cache.Tasks, nil
}
