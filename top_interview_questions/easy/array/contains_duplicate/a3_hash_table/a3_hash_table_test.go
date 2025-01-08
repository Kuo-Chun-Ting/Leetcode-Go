package a3_hash_table

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenHasDuplicateNum(t *testing.T) {
	// Arrange
	nums := []int{1, 2, 3, 1}

	// Act
	result := containsDuplicate(nums)

	// Assert
	assert.Equal(t, true, result)
}

func TestWhenNoDuplicateNum(t *testing.T) {
	// Arrange
	nums := []int{1, 2, 3, 4}

	// Act
	result := containsDuplicate(nums)

	// Assert
	assert.Equal(t, false, result)
}

func TestWhenMultipleDuplicateNumPair(t *testing.T) {
	// Arrange
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}

	// Act
	result := containsDuplicate(nums)

	// Assert
	assert.Equal(t, true, result)
}
