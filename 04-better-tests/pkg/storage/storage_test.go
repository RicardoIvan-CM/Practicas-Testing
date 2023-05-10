package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
func (s *storage) GetValue(key string) interface{} {
	file, err := os.ReadFile(s.file)
	if err != nil {
		panic(err)
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
	file := "../config.json"
	return &storage{file: file}
}
*/

func TestNewStorage(t *testing.T) {
	//Arrange
	var expected = &storage{
		file: "config.json",
	}

	//Act
	s := NewStorage()

	//Assert
	assert.Equal(t, expected, s)
}

func TestNewStorageFromURL(t *testing.T) {
	//Arrange
	var expected = &storage{
		file: "config.json",
	}

	//Act
	s := NewStorageFromURL("config.json")

	//Assert
	assert.Equal(t, expected, s)
}

func TestGetValue(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		//Arrange
		var storage = NewStorage()
		var expected float64 = 100

		//Act
		result := storage.GetValue("tuna_speed")

		//Assert
		assert.Equal(t, expected, result)
	})

	t.Run("Not Found", func(t *testing.T) {
		//Arrange
		var storage = NewStorageFromURL("lol.json")

		//Act
		result := storage.GetValue("tuna_speed")

		//Assert
		assert.Nil(t, result)
	})

	t.Run("Malformed JSON", func(t *testing.T) {
		//Arrange
		var storage = NewStorageFromURL("config_bad.json")

		//Act
		result := storage.GetValue("tuna_speed")

		//Assert
		assert.Nil(t, result)
	})

	t.Run("Property Not Found", func(t *testing.T) {
		//Arrange
		var storage = NewStorageFromURL("config.json")

		//Act
		result := storage.GetValue("tuna_type")

		//Assert
		assert.Nil(t, result)
	})
}
