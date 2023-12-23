package a1_brute_force

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenTwoNumAreDifferent(t *testing.T) {
	// Arrange
	target := 9
	nums := []int{1, 6, 6, 8, 4}

	// Act
	result := twoSum(nums, target)

	// Assert
	expectedResult := []int{0, 3}
	assert.Equal(t, expectedResult, result)
}

func TestWhenTwoNumAreTheSame(t *testing.T) {
	// Arrange
	target := 10
	nums := []int{1, 6, 6, 8, 5, 7, 7, 5}

	// Act
	result := twoSum(nums, target)

	// Assert
	expectedResult := []int{4, 7}
	assert.Equal(t, expectedResult, result)
}
