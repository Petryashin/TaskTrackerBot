package memory

import (
	"encoding/json"
	"io/ioutil"
)

type commonCache map[int64]Cache

type Cache struct {
	Tasks []Task
}

func (c *commonCache) mustPutCache() error {
	marshal, err := json.Marshal(c)

	if err != nil {
		return err
	}
	filename := "cache.json"

	return ioutil.WriteFile(filename, marshal, 0600)

}

func mustParseCache() (*commonCache, error) {
	filename := "cache.json"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return &commonCache{}, err
	}
	c := &commonCache{}
	if err := json.Unmarshal(body, c); err != nil {
		return &commonCache{}, err
	}
	return c, nil
}
