package shark

import (
	"fmt"
	"integrationtests/pkg/storage"
	"integrationtests/prey"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockStorage struct {
	speed     float64
	x         float64
	y         float64
	tunaSpeed float64
}

func (s *mockStorage) GetValue(key string) interface{} {
	switch key {
	case "white_shark_speed":
		return s.speed
	case "white_shark_x":
		return s.x
	case "white_shark_y":
		return s.y
	case "tuna_speed":
		return s.tunaSpeed
	default:
		return nil
	}
}

func NewMockStorage(speed float64, x float64, y float64, tunaSpeed float64) storage.Storage {
	return &mockStorage{
		speed,
		x,
		y,
		tunaSpeed,
	}
}

type catchSimulatorMock struct {
	// max time to catch the prey in seconds
	maxTimeToCatch float64
	GLDSpy         bool
}

func (r *catchSimulatorMock) CanCatch(distance float64, speed float64, catchSpeed float64) bool {
	timeToCatch := distance / (speed - catchSpeed)
	return timeToCatch > 0 && timeToCatch <= r.maxTimeToCatch
}

func (r *catchSimulatorMock) GetLinearDistance(position [2]float64) float64 {
	r.GLDSpy = true
	x := big.NewFloat(position[0])
	y := big.NewFloat(position[1])
	z := x.Add(x.Mul(x, x), y.Mul(y, y))
	res, _ := z.Sqrt(z).Float64()
	return res
}

func NewCatchSimulatorMock(maxTimeToCatch float64) catchSimulatorMock {
	return catchSimulatorMock{
		maxTimeToCatch: maxTimeToCatch,
		GLDSpy:         false,
	}
}

func TestSharkHuntsSuccessfully(t *testing.T) {
	//Arrange
	var cs = NewCatchSimulatorMock(50)
	var storage = NewMockStorage(200, 100, 100, 100)
	var shark = CreateWhiteShark(&cs, storage)
	var tuna = prey.CreateTuna(storage)

	//Act
	err := shark.Hunt(tuna)
	//Assert
	assert.NoError(t, err)
	assert.Equal(t, true, cs.GLDSpy)
}

func TestSharkCannotCatchSpeed(t *testing.T) {
	//Arrange
	var cs = NewCatchSimulatorMock(50)
	var storage = NewMockStorage(100, 100, 100, 200)
	var shark = CreateWhiteShark(&cs, storage)
	var tuna = prey.CreateTuna(storage)
	var expected error = fmt.Errorf("could not hunt the prey")

	//Act
	err := shark.Hunt(tuna)
	//Assert
	assert.EqualError(t, err, expected.Error())
	assert.Equal(t, true, cs.GLDSpy)
}

func TestSharkCannotCatchDistance(t *testing.T) {
	//Arrange
	var cs = NewCatchSimulatorMock(5)
	var storage = NewMockStorage(250, 330, 330, 200)
	var shark = CreateWhiteShark(&cs, storage)
	var tuna = prey.CreateTuna(storage)

	var expected error = fmt.Errorf("could not hunt the prey")

	//Act
	err := shark.Hunt(tuna)
	//Assert
	assert.EqualError(t, err, expected.Error())
	assert.Equal(t, true, cs.GLDSpy)
}
