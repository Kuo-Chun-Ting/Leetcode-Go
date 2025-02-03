package a1_two_indexes_approach

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhen1ThenReturn1(t *testing.T) {
	// Arrange
	sortedNums := []int{1}

	// Act
	result := removeDuplicates(sortedNums)

	// Assert
	expectedArray := []int{1}
	expectedLen := len(expectedArray)
	assert.Equal(t, expectedLen, result)
	assert.Equal(t, expectedArray, sortedNums[:expectedLen])
}

func TestWhen123ThenReturn123(t *testing.T) {
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

func TestWhen1223ThenReturn123(t *testing.T) {
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

func TestWhen112ThenReturn12(t *testing.T) {
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
