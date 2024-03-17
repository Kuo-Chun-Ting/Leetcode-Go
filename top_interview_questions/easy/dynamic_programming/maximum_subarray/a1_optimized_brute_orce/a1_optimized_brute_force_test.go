package a1_optimized_brute_force

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenOnlyOneNumberThenReturnTheNumber(t *testing.T) {
	// Arrange
	arr := []int{1}

	// Act
	maxSubArrSum := maxSubArray(arr)

	// Assert
	expectedMaxSubArrSum := 1
	assert.Equal(t, expectedMaxSubArrSum, maxSubArrSum)
}

func TestWhenHasMaxSubArrSum(t *testing.T) {
	// Arrange
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	// Act
	maxSubArrSum := maxSubArray(arr)

	// Assert
	expectedMaxSubArrSum := 6
	assert.Equal(t, expectedMaxSubArrSum, maxSubArrSum)
}

func TestWhenTheEntireArrIsTheAnswer(t *testing.T) {
	// Arrange
	arr := []int{5, 4, -1, 7, 8}

	// Act
	maxSubArrSum := maxSubArray(arr)

	// Assert
	expectedMaxSubArrSum := 23
	assert.Equal(t, expectedMaxSubArrSum, maxSubArrSum)
}
