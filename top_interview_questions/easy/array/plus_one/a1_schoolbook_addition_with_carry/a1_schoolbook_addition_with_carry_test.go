package a1_schoolbook_addition_with_carry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusOneWhenZeroReturnOne(t *testing.T) {
	// Arrange
	digits := []int{0}

	// Act
	result := plusOne(digits)

	// Assert
	assert.Equal(t, []int{1}, result)
}

func TestPlusOneWhenLastDigitNoOverflow(t *testing.T) {
	// Arrange
	digits := []int{1, 1}

	// Act
	result := plusOne(digits)

	// Assert
	assert.Equal(t, []int{1, 2}, result)
}

func TestPlusOneWhenLastDigitOverflow(t *testing.T) {
	// Arrange
	digits := []int{1, 9}

	// Act
	result := plusOne(digits)

	// Assert
	assert.Equal(t, []int{2, 0}, result)
}

func TestPlusOneWhenAllDigitsOverflow(t *testing.T) {
	// Arrange
	digits := []int{9, 9}

	// Act
	result := plusOne(digits)

	// Assert
	assert.Equal(t, []int{1, 0, 0}, result)
}
