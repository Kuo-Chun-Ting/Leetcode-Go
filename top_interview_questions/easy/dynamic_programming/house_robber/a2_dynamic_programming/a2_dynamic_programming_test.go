package a1_recursion_with_memoization

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenOnlyOneNumberThenReturnTheNumber(t *testing.T) {
	// Arrange
	arr := []int{1}

	// Act
	sum := rob(arr)

	// Assert
	expectedSum := 1
	assert.Equal(t, expectedSum, sum)
}

func TestWhenMoreThanOneNumberThenReturnMaxAmount(t *testing.T) {
	// Arrange
	arr := []int{1, 2, 3, 4}

	// Act
	sum := rob(arr)

	// Assert
	expectedSum := 6
	assert.Equal(t, expectedSum, sum)
}
