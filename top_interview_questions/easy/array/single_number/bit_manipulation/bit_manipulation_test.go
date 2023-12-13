package bit_manipulation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumberWhenOneElementThenReturnIt(t *testing.T) {
	// Arrange
	singleNum := 0
	nums := []int{singleNum}

	// Act
	result := singleNumber(nums)

	// Assert
	expectedElement := 0
	assert.Equal(t, expectedElement, result)
}

func TestSingleNumberWhenMoreThanOneElementsThenReturnSingleOne(t *testing.T) {
	// Arrange
	singleNum := 1
	nums := []int{2, singleNum, 3, 2, 3, 11, 11}

	// Act
	result := singleNumber(nums)

	// Assert
	expectedElement := 1
	assert.Equal(t, expectedElement, result)
}
