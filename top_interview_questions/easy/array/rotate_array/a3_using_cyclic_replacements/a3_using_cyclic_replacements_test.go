package a3_using_cyclic_replacements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateWhenOneItems(t *testing.T) {
	// Arrange
	nums := []int{1}
	k := 3

	// Act
	rotate(nums, k)

	// Assert
	expectedNums := []int{1}
	assert.Equal(t, expectedNums, nums)
}

func TestRotateWhenKLessThanLengthOfNums(t *testing.T) {
	// Arrange
	nums := []int{1, 2, 3, 4, 5}
	k := 3

	// Act
	rotate(nums, k)

	// Assert
	expectedNums := []int{3, 4, 5, 1, 2}
	assert.Equal(t, expectedNums, nums)
}

func TestRotateWhenKEqualToLengthOfNums(t *testing.T) {
	// Arrange
	nums := []int{1, 2, 3, 4, 5}
	k := 5

	// Act
	rotate(nums, k)

	// Assert
	expectedNums := []int{1, 2, 3, 4, 5}
	assert.Equal(t, expectedNums, nums)
}

func TestRotateWhenKLargerThanLengthOfNums(t *testing.T) {
	// Arrange
	nums := []int{1, 2, 3, 4, 5}
	k := 6

	// Act
	rotate(nums, k)

	// Assert
	expectedNums := []int{5, 1, 2, 3, 4}
	assert.Equal(t, expectedNums, nums)
}

func TestRotateWhenNModKNot0(t *testing.T) {
	// Arrange
	nums := []int{1, 2, 3, 4, 5}
	k := 3

	// Act
	rotate(nums, k)

	// Assert
	expectedNums := []int{3, 4, 5, 1, 2}
	assert.Equal(t, expectedNums, nums)
}

func TestRotateWhenNModKIs0(t *testing.T) {
	// Arrange
	nums := []int{1, 2, 3, 4, 5, 6}
	k := 3

	// Act
	rotate(nums, k)

	// Assert
	expectedNums := []int{4, 5, 6, 1, 2, 3}
	assert.Equal(t, expectedNums, nums)
}
