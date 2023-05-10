package prey

import (
	"integrationtests/pkg/storage"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
type tuna struct {
	maxSpeed float64
}

func (t *tuna) GetSpeed() float64 {
	return t.maxSpeed
}

func CreateTuna(storage storage.Storage) Prey {
	return &tuna{
		maxSpeed: storage.GetValue("tuna_speed").(float64),
	}
}
*/

func TestCreateTuna(t *testing.T) {
	//Arrange
	var storage = storage.NewStorageFromURL("../pkg/storage/config.json")
	var expected = &tuna{
		maxSpeed: 100,
	}

	//Act
	prey := CreateTuna(storage)
	//Assert
	assert.Equal(t, expected, prey)
}

func TestGetSpeed(t *testing.T) {
	//Arrange
	var tuna = &tuna{
		maxSpeed: 100,
	}
	var expected float64 = 100

	//Act
	result := tuna.GetSpeed()
	//Assert
	assert.Equal(t, expected, result)
}
