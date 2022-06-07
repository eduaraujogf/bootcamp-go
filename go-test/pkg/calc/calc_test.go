package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	result := Sum(5, 5)
	expectedResult := 10

	if expectedResult == result {
		t.FailNow()
	}

	assert.Equal(t, expectedResult, result, "a soma está errada")
}

func TestSub(t *testing.T) {
	result := Sub(10, 5)
	expectedResult := 5

	assert.Equal(t, expectedResult, result, "a subtração está errada")
}

func TestMultiply(t *testing.T) {
	result := Multiply(10, 5)
	expectedResult := 50

	assert.Equal(t, expectedResult, result, "a multiplicação está errada")
}
func TestDivision(t *testing.T) {
	result := Division(10, 5)
	expectedResult := 2

	assert.Equal(t, expectedResult, result, "a divisão está errada")
}
