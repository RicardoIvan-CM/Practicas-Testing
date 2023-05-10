package simulator

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCatchSimulator(t *testing.T) {
	//Arrange
	var expected = &catchSimulator{
		maxTimeToCatch: 50,
	}
	//Act
	var result = NewCatchSimulator(50)
	//Assert
	assert.Equal(t, expected, result)
}

func TestCanCatch(t *testing.T) {
	t.Run("Could Not Catch", func(t *testing.T) {
		//Arrange
		var simulator = NewCatchSimulator(100)
		var expected = false
		//Act
		var result = simulator.CanCatch(50, 100, 300)
		//Assert
		assert.Equal(t, expected, result)
	})

	t.Run("Could Catch", func(t *testing.T) {
		//Arrange
		var simulator = NewCatchSimulator(10)
		var expected = true
		//Act
		var result = simulator.CanCatch(50, 300, 100)
		//Assert
		assert.Equal(t, expected, result)
	})
}

func TestGetLinearDistance(t *testing.T) {
	//Arrange
	var positions = [2]float64{2, 2}
	var simulator = NewCatchSimulator(10)
	var expected = math.Sqrt(8)
	//Act
	var result = simulator.GetLinearDistance(positions)
	//Assert
	assert.Equal(t, expected, result)
}
