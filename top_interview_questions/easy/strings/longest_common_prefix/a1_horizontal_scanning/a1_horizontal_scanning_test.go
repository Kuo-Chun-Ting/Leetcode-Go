package a1_horizontal_scanning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenStrLenIsZeroThenReturnEmptyString(t *testing.T) {
	// Arrange
	str := []string{}

	// Act
	result := longestCommonPrefix(str)

	// Assert
	expectedResult := ""
	assert.Equal(t, expectedResult, result)
}

func TestWhenLCPNotExsitingThenReturnLCP(t *testing.T) {
	// Arrange
	str := []string{}

	// Act
	result := longestCommonPrefix(str)

	// Assert
	expectedResult := ""
	assert.Equal(t, expectedResult, result)
}

func TestWhenLCPExsitingThenReturnLCP(t *testing.T) {
	// Arrange
	str := []string{"flower", "flow", "flight"}

	// Act
	result := longestCommonPrefix(str)

	// Assert
	expectedResult := "fl"
	assert.Equal(t, expectedResult, result)
}
