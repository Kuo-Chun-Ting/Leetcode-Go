package a2_frequency_counter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenOneNonAnagramCharReturnTrue(t *testing.T) {
	// Arrange
	s1 := "a"
	s2 := "b"

	// Act
	result := isAnagram(s1, s2)

	// Assert
	assert.Equal(t, false, result)
}

func TestWhenOneAnagramCharReturnTrue(t *testing.T) {
	// Arrange
	s1 := "a"
	s2 := "a"

	// Act
	result := isAnagram(s1, s2)

	// Assert
	assert.Equal(t, true, result)
}

func TestWhenStringLengthDifferentThenReturnFalse(t *testing.T) {
	// Arrange
	s1 := "a"
	s2 := "ab"

	// Act
	result := isAnagram(s1, s2)

	// Assert
	assert.Equal(t, false, result)
}

func TestWhenAnagramCharReturnTrue1(t *testing.T) {
	// Arrange
	s1 := "anagram"
	s2 := "nagaram"

	// Act
	result := isAnagram(s1, s2)

	// Assert
	assert.Equal(t, true, result)
}

func TestWhenNonAnagramCharReturnTrue1(t *testing.T) {
	// Arrange
	s1 := "rat"
	s2 := "car"

	// Act
	result := isAnagram(s1, s2)

	// Assert
	assert.Equal(t, false, result)
}
