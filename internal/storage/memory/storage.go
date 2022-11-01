package memory

import (
	"encoding/json"
	"io/ioutil"
)

type commonCache map[int]Cache

type Cache struct {
	Tasks []Task
}

func (c *Cache) mustPutCache() error {
	marshal, err := json.Marshal(c)

	if err != nil {
		return err
	}
	filename := "cache.json"

	return ioutil.WriteFile(filename, marshal, 0600)

}

func mustParseCache() (*Cache, error) {
	filename := "cache.json"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return &Cache{}, err
	}
	c := &Cache{}
	if err := json.Unmarshal(body, c); err != nil {
		return &Cache{}, err
	}
	return c, nil
}
