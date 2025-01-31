package a1_loop_iteration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenNegativeNumber(t *testing.T) {
	// Arrange
	n := -1

	// Act
	result := isPowerOfThree(n)

	// Assert
	assert.Equal(t, false, result)
}

func TestWhen0(t *testing.T) {
	// Arrange
	n := 0

	// Act
	result := isPowerOfThree(n)

	// Assert
	assert.Equal(t, false, result)
}

func TestWhen1(t *testing.T) {
	// Arrange
	n := 1

	// Act
	result := isPowerOfThree(n)

	// Assert
	assert.Equal(t, true, result)
}

func TestWhen2(t *testing.T) {
	// Arrange
	n := 2

	// Act
	result := isPowerOfThree(n)

	// Assert
	assert.Equal(t, false, result)
}

func TestWhen3(t *testing.T) {
	// Arrange
	n := 3

	// Act
	result := isPowerOfThree(n)

	// Assert
	assert.Equal(t, true, result)
}

func TestWhen4(t *testing.T) {
	// Arrange
	n := 4

	// Act
	result := isPowerOfThree(n)

	// Assert
	assert.Equal(t, false, result)
}

func TestWhen9(t *testing.T) {
	// Arrange
	n := 9

	// Act
	result := isPowerOfThree(n)

	// Assert
	assert.Equal(t, true, result)
}

func TestWhen27(t *testing.T) {
	// Arrange
	n := 27

	// Act
	result := isPowerOfThree(n)

	// Assert
	assert.Equal(t, true, result)
}
