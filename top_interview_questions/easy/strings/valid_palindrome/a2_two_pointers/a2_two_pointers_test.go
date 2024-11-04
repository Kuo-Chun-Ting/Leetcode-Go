package a2_two_pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenEmptyStringThenReturnTrue(t *testing.T) {
	// Arrange
	str := " "

	// Act
	result := isPalindrome(str)

	// Assert
	assert.Equal(t, true, result)
}

func TestWhenOneCharThenReturnTrue(t *testing.T) {
	// Arrange
	str := "a"

	// Act
	result := isPalindrome(str)

	// Assert
	assert.Equal(t, true, result)
}

func TestWhenPalindromeThenReturnTrue(t *testing.T) {
	// Arrange
	str := "A man, a plan, a canal: Panama"

	// Act
	result := isPalindrome(str)

	// Assert
	assert.Equal(t, true, result)
}

func TestWhenNotPalindromeThenReturnTrue(t *testing.T) {
	// Arrange
	str := "race a car"

	// Act
	result := isPalindrome(str)

	// Assert
	assert.Equal(t, false, result)
}
