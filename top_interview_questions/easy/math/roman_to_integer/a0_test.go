package a0

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenIII(t *testing.T) {
	// Arrange
	n := "III"

	// Act
	result := romanToInt(n)

	// Assert
	assert.Equal(t, 3, result)
}

func TestWhenLVIII(t *testing.T) {
	// Arrange
	n := "LVIII"

	// Act
	result := romanToInt(n)

	// Assert
	assert.Equal(t, 58, result)
}

func TestWhenMCMXCIV(t *testing.T) {
	// Arrange
	n := "MCMXCIV"

	// Act
	result := romanToInt(n)

	// Assert
	assert.Equal(t, 1994, result)
}

func TestWhenIV(t *testing.T) {
	// Arrange
	n := "IV"

	// Act
	result := romanToInt(n)

	// Assert
	assert.Equal(t, 4, result)
}

func TestWhenIX(t *testing.T) {
	// Arrange
	n := "IX"

	// Act
	result := romanToInt(n)

	// Assert
	assert.Equal(t, 9, result)
}
