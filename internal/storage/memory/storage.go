package memory

import (
	"encoding/json"
	"io/ioutil"
)

type Cache struct {
	Tasks []Task
}

func (c *Cache) mustPutCache() error {
	marshal, err := json.Marshal(c)

	if err != nil {
		panic(err)
	}
	filename := "cache.json"

	return ioutil.WriteFile(filename, marshal, 0600)

}

func mustParseCache() *Cache {
	filename := "cache.json"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return &Cache{}
	}
	c := &Cache{}
	if err := json.Unmarshal(body, c); err != nil {
		panic(err)
	}
	return c
}
