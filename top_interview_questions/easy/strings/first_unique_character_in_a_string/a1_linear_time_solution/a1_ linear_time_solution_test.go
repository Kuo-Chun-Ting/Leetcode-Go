package a1_linear_time_solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenUniqueCharAtFirst(t *testing.T) {
	// Arrange
	str := "abc"

	// Act
	result := firstUniqChar(str)

	// Assert
	expectedIdx := 0
	assert.Equal(t, expectedIdx, result)
}

func TestWhenUniqueCharAtMiddle(t *testing.T) {
	// Arrange
	str := "acabc"

	// Act
	result := firstUniqChar(str)

	// Assert
	expectedIdx := 3
	assert.Equal(t, expectedIdx, result)
}

func TestWhenUniqueCharAtEnd(t *testing.T) {
	// Arrange
	str := "ababc"

	// Act
	result := firstUniqChar(str)

	// Assert
	expectedIdx := 4
	assert.Equal(t, expectedIdx, result)
}

func TestWhenNoUniqueChar(t *testing.T) {
	// Arrange
	str := "aabbcc"

	// Act
	result := firstUniqChar(str)

	// Assert
	expectedIdx := -1
	assert.Equal(t, expectedIdx, result)
}

func Test(t *testing.T) {
	// Arrange
	str := "loveleetcode"

	// Act
	result := firstUniqChar(str)

	// Assert
	expectedIdx := 2
	assert.Equal(t, expectedIdx, result)
}
