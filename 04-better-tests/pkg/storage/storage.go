package storage

import (
	"encoding/json"
	"os"
)

type Storage interface {
	GetValue(key string) interface{}
}

type storage struct {
	file string
}

func (s *storage) GetValue(key string) interface{} {
	file, err := os.ReadFile(s.file)
	if err != nil {
		//panic(err)
		return nil
	}

	data := make(map[string]interface{})
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil
	}

	if v, ok := data[key]; ok {
		return v
	}

	return nil
}

func NewStorage() Storage {
	url := "config.json"
	return &storage{file: url}
}

func NewStorageFromURL(url string) Storage {
	return &storage{file: url}
}
