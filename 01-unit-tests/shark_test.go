package hunt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSharkHuntsSuccessfully(t *testing.T) {
	//Arrange
	var s Shark = Shark{
		hungry: true,
		tired:  false,
		speed:  2,
	}
	var p Prey = Prey{
		name:  "Fishy",
		speed: 1,
	}

	//Act
	err := s.Hunt(&p)

	//Assert
	assert.NoError(t, err)
	assert.Equal(t, false, s.hungry)
	assert.Equal(t, true, s.tired)
}

func TestSharkCannotHuntBecauseIsTired(t *testing.T) {
	//Arrange
	var s Shark = Shark{
		hungry: true,
		tired:  true,
		speed:  0,
	}
	var p Prey = Prey{
		name:  "Fishy",
		speed: 1,
	}
	var expected error = fmt.Errorf("cannot hunt, i am really tired")

	//Act
	err := s.Hunt(&p)

	//Assert
	assert.EqualError(t, err, expected.Error())
}

func TestSharkCannotHuntBecaisIsNotHungry(t *testing.T) {
	//Arrange
	var s Shark = Shark{
		hungry: false,
		tired:  false,
		speed:  0,
	}
	var p Prey = Prey{
		name:  "Fishy",
		speed: 1,
	}
	var expected error = fmt.Errorf("cannot hunt, i am not hungry")

	//Act
	err := s.Hunt(&p)

	//Assert
	assert.EqualError(t, err, expected.Error())
}

func TestSharkCannotReachThePrey(t *testing.T) {
	//Arrange
	var s Shark = Shark{
		hungry: true,
		tired:  false,
		speed:  0,
	}
	var p Prey = Prey{
		name:  "Fishy",
		speed: 1,
	}
	var expected error = fmt.Errorf("could not catch it")

	//Act
	err := s.Hunt(&p)

	//Assert
	assert.EqualError(t, err, expected.Error())
}

/*
func TestSharkHuntNilPrey(t *testing.T) {
	//Arrange
	var s Shark = Shark{
		hungry: true,
		tired:  false,
		speed:  2,
	}

	//Act

	//Assert
	assert.Panics(t, func() {
		s.Hunt(nil)
	})
}*/

func TestSharkHuntNilPrey(t *testing.T) {
	//Arrange
	var s Shark = Shark{
		hungry: true,
		tired:  false,
		speed:  2,
	}

	var expected error = fmt.Errorf("prey must not be nil")

	//Act
	err := s.Hunt(nil)

	//Assert
	assert.EqualError(t, err, expected.Error())
}
