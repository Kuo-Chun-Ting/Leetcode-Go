package a2_two_pass_hash_table

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

func TestWhenSameIndexCanGetTheTargetThenShouldNotUseTheIndex(t *testing.T) {
	// Arrange
	target := 10
	// If the logic can't handle the index well(i != j) it will return {0, 0} where 5 + 5 = 10
	// In this case it should return {1, 6}
	nums := []int{5, 6, 9, 8, 7, 7, 4}

	// Act
	result := twoSum(nums, target)

	// Assert
	expectedResult := []int{1, 6}
	assert.Equal(t, expectedResult, result)
}
