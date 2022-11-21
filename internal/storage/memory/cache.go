package memory

import "errors"

type Storage struct {
	Cache *commonCache
}

type Task struct {
	Text string
}

func NewStorage() (Storage, error) {
	tasks, err := mustParseCache()
	if err != nil {
		return Storage{Cache: &commonCache{}}, err
	}

	return Storage{Cache: tasks}, nil
}

func (c Storage) Add(chatId int64, message string) error {
	if _, ok := (*c.Cache)[chatId]; !ok {
		(*c.Cache)[chatId] = Cache{}
	}
	if entry, ok := (*c.Cache)[chatId]; ok {
		entry.Tasks = append(entry.Tasks, Task{message})
		(*c.Cache)[chatId] = entry
		c.Cache.mustPutCache()
		return nil
	}
	return errors.New("Can't add field")
}

func (c Storage) Remove(chatId int64, key int) error {
	if entry, ok := (*c.Cache)[chatId]; ok {
		if key < 1 || key > len(entry.Tasks) {
			return errors.New("Задачи с таким номером не существует")
		}
		entry.Tasks = append(entry.Tasks[:key-1], entry.Tasks[key:]...)
		(*c.Cache)[chatId] = entry
		c.Cache.mustPutCache()
		return nil
	}
	return nil
}

func (c Storage) List(chatId int64) ([]Task, error) {
	return (*c.Cache)[chatId].Tasks, nil
}
