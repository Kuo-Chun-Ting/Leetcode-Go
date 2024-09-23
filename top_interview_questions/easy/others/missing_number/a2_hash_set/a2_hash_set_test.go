package a2_hash_set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenMissingValueAtFirst(t *testing.T) {
	// Arrange
	nums := []int{1, 2}

	// Act
	result := missingNumber(nums)

	// Assert
	assert.Equal(t, 0, result)
}

func TestWhenMissingValueAtLast(t *testing.T) {
	// Arrange
	nums := []int{0, 1}

	// Act
	result := missingNumber(nums)

	// Assert
	assert.Equal(t, 2, result)
}

func TestWhenMissingValueInMiddle(t *testing.T) {
	// Arrange
	nums := []int{0, 2}

	// Act
	result := missingNumber(nums)

	// Assert
	assert.Equal(t, 1, result)
}
