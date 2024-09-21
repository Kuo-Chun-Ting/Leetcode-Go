package a3_brian_kernighans_algorithm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhen1And4(t *testing.T) {
	// Arrange
	x := 1 // 0001
	y := 4 // 0100

	// Act
	result := hammingDistance(x, y)

	// Assert
	assert.Equal(t, 2, result)
}

func TestWhen1And5(t *testing.T) {
	// Arrange
	x := 1 // 0001
	y := 5 // 0101

	// Act
	result := hammingDistance(x, y)

	// Assert
	assert.Equal(t, 1, result)
}
