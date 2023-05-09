package shark

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

type PreyStub struct {
	maxSpeed float64
}

func (ps *PreyStub) GetSpeed() float64 {
	return ps.maxSpeed
}

var ps1 = PreyStub{100}
var ps2 = PreyStub{200}

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
	var shark = CreateWhiteShark(&cs)

	//Act
	err := shark.Hunt(&ps1)
	//Assert
	assert.NoError(t, err)
	assert.Equal(t, true, cs.GLDSpy)
}

func TestSharkCannotCatchSpeed(t *testing.T) {
	//Arrange
	var cs = NewCatchSimulatorMock(50)
	var shark = CreateWhiteShark(&cs)
	var expected error = fmt.Errorf("could not hunt the prey")

	//Act
	err := shark.Hunt(&ps2)
	//Assert
	assert.EqualError(t, err, expected.Error())
	assert.Equal(t, true, cs.GLDSpy)
}

func TestSharkCannotCatchDistance(t *testing.T) {
	//Arrange
	var cs = NewCatchSimulatorMock(5)
	var shark = whiteShark{
		speed:     250,
		position:  [2]float64{330, 330},
		simulator: &cs,
	}

	var expected error = fmt.Errorf("could not hunt the prey")

	//Act
	err := shark.Hunt(&ps2)
	//Assert
	assert.EqualError(t, err, expected.Error())
	assert.Equal(t, true, cs.GLDSpy)
}
