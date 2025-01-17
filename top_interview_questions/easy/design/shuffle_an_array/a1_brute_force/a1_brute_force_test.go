package a1_brute_force

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenConstructorCalledThenSolutionInitialized(t *testing.T) {
	// Arrange
	nums := []int{1, 2, 3}

	// Act
	solution := Constructor(nums)

	// Assert
	assert.Equal(t, nums, solution.Reset())
}

func TestWhenResetCalledThenOriginalArrayReturned(t *testing.T) {
	// Arrange
	nums := []int{1, 2, 3}
	solution := Constructor(nums)

	// Act
	result := solution.Reset()

	// Assert
	expectedInitNums := []int{1, 2, 3}
	assert.Equal(t, expectedInitNums, result)
}

func TestWhenShuffleCalledThenArrayIsShuffled(t *testing.T) {
	// Arrange
	nums := []int{1, 2, 3}
	solution := Constructor(nums)

	// Act
	shuffled := solution.Shuffle()

	// Assert
	assert.ElementsMatch(t, nums, shuffled)
}
