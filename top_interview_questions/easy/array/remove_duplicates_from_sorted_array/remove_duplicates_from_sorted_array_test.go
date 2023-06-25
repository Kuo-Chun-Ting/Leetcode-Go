package remove_duplicates_from_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicatesWhenNumsIsNilThenReturn0(t *testing.T) {
	// Arrange
	var sortedNums []int

	// Act
	result := removeDuplicates(sortedNums)

	// Assert
	assert.Equal(t, 0, result)
}

func TestRemoveDuplicatesWhenNumsEmptyThenReturn0(t *testing.T) {
	// Arrange
	sortedNums := []int{}

	// Act
	result := removeDuplicates(sortedNums)

	// Assert
	assert.Equal(t, 0, result)
}

func TestRemoveDuplicatesWhenNumsHasNoDuplicateThenReturnExpectedLen(t *testing.T) {
	// Arrange
	sortedNums := []int{1, 2, 3}

	// Act
	result := removeDuplicates(sortedNums)

	// Assert
	expectedArray := []int{1, 2, 3}
	expectedLen := len(expectedArray)
	assert.Equal(t, expectedLen, result)
	assert.Equal(t, expectedArray, sortedNums[:expectedLen])
}

func TestRemoveDuplicatesWhenNumsHasDuplicateThenReturnExpectedLen(t *testing.T) {
	// Arrange
	sortedNums := []int{1, 2, 2, 3}

	// Act
	result := removeDuplicates(sortedNums)

	// Assert
	expectedArray := []int{1, 2, 3}
	expectedLen := len(expectedArray)
	assert.Equal(t, expectedLen, result)
	assert.Equal(t, expectedArray, sortedNums[:expectedLen])
}

func TestRemoveDuplicatesWhenNumsHasDuplicateThenReturnExpectedLenTest(t *testing.T) {
	// Arrange
	sortedNums := []int{1, 1, 2}

	// Act
	result := removeDuplicates(sortedNums)

	// Assert
	expectedArray := []int{1, 2}
	expectedLen := len(expectedArray)
	assert.Equal(t, expectedLen, result)
	assert.Equal(t, expectedArray, sortedNums[:expectedLen])
}
