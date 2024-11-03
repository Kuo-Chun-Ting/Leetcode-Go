package a1_pop_and_push_digits_and_check_before_overflow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenOnePositiveNumber(t *testing.T) {
	// Arrange
	i := 1

	// Act
	result := reverse(i)

	// Assert
	expectedResult := 1
	assert.Equal(t, expectedResult, result)
}

func TestWhenOneNegativeNumber(t *testing.T) {
	// Arrange
	i := -1

	// Act
	result := reverse(i)

	// Assert
	expectedResult := -1
	assert.Equal(t, expectedResult, result)
}

func TestWhenPositiveNumberAndNotOverflow(t *testing.T) {
	// Arrange
	i := 123

	// Act
	result := reverse(i)

	// Assert
	expectedResult := 321
	assert.Equal(t, expectedResult, result)
}

func TestWhenPositiveNumberAndOverflow(t *testing.T) {
	// Arrange
	i := 1234567899 // max integer is 2147483647

	// Act
	result := reverse(i)

	// Assert
	expectedResult := 0
	assert.Equal(t, expectedResult, result)
}

func TestWhenNegativeNumberAndNotOverflow(t *testing.T) {
	// Arrange
	i := -123

	// Act
	result := reverse(i)

	// Assert
	expectedResult := -321
	assert.Equal(t, expectedResult, result)
}

func TestWhenNegativeNumberAndOverflow(t *testing.T) {
	// Arrange
	i := -1234567899 // min integer is -2147483648

	// Act
	result := reverse(i)

	// Assert
	expectedResult := 0
	assert.Equal(t, expectedResult, result)
}
