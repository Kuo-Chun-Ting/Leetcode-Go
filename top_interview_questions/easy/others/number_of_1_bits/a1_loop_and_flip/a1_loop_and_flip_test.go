package a1_loop_and_flip

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenZero(t *testing.T) {
	// Arrange
	n := 0

	// Act
	result := hammingWeight(n)

	// Assert
	assert.Equal(t, 0, result)
}

func TestWhenOne(t *testing.T) {
	// Arrange
	n := 1

	// Act
	result := hammingWeight(n)

	// Assert
	assert.Equal(t, 1, result)
}

func TestWhenThreeBits(t *testing.T) {
	// Arrange
	n := 7

	// Act
	result := hammingWeight(n)

	// Assert
	assert.Equal(t, 3, result)
}
