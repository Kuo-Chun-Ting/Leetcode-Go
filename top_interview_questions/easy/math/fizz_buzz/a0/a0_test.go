package a0

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhen3(t *testing.T) {
	// Arrange
	n := 3

	// Act
	result := fizzBuzz(n)

	// Assert
	expectedResult := []string{"1", "2", "Fizz"}
	assert.Equal(t, expectedResult, result)
}

func TestWhen5(t *testing.T) {
	// Arrange
	n := 5

	// Act
	result := fizzBuzz(n)

	// Assert
	expectedResult := []string{"1", "2", "Fizz", "4", "Buzz"}
	assert.Equal(t, expectedResult, result)
}

func TestWhen15(t *testing.T) {
	// Arrange
	n := 15

	// Act
	result := fizzBuzz(n)

	// Assert
	expectedResult := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	assert.Equal(t, expectedResult, result)
}
