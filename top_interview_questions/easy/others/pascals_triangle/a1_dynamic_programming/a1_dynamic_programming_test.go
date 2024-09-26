package a1_dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenInputOne(t *testing.T) {
	// Arrange
	input := 1
	// Act
	result := generate(input)

	// Assert
	expectedOutput := [][]int{{1}}
	assert.Equal(t, expectedOutput, result)
}

func TestWhenInputTwo(t *testing.T) {
	// Arrange
	input := 2
	// Act
	result := generate(input)

	// Assert
	expectedOutput := [][]int{{1}, {1, 1}}
	assert.Equal(t, expectedOutput, result)
}

func TestWhenInputThree(t *testing.T) {
	// Arrange
	input := 3
	// Act
	result := generate(input)

	// Assert
	expectedOutput := [][]int{{1}, {1, 1}, {1, 2, 1}}
	assert.Equal(t, expectedOutput, result)
}
