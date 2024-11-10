package me_check_every_needle_first_char

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenNeedleLengthGreaterThanHaystackThenReturnNegativeOne(t *testing.T) {
	// Arrange
	haystack := "abc"
	needle := "abcdef"

	// Act
	result := strStr(haystack, needle)

	// Assert
	expectedResult := -1
	assert.Equal(t, expectedResult, result)
}

func TestWhenNoOccurenceThenReturnNegativeOne(t *testing.T) {
	// Arrange
	haystack := "abc"
	needle := "defg"

	// Act
	result := strStr(haystack, needle)

	// Assert
	expectedResult := -1
	assert.Equal(t, expectedResult, result)
}

func TestWhenOneOccurenceAtTheBeginingThenReturnTheOccurrenceIndex(t *testing.T) {
	// Arrange
	haystack := "abcdef"
	needle := "abc"

	// Act
	result := strStr(haystack, needle)

	// Assert
	expectedResult := 0
	assert.Equal(t, expectedResult, result)
}

func TestWhenOneOccurenceAtTheMiddleThenReturnTheOccurrenceIndex(t *testing.T) {
	// Arrange
	haystack := "dabcef"
	needle := "abc"

	// Act
	result := strStr(haystack, needle)

	// Assert
	expectedResult := 1
	assert.Equal(t, expectedResult, result)
}

func TestWhenOneOccurenceAtTheEndThenReturnTheOccurrenceIndex(t *testing.T) {
	// Arrange
	haystack := "defabc"
	needle := "abc"

	// Act
	result := strStr(haystack, needle)

	// Assert
	expectedResult := 3
	assert.Equal(t, expectedResult, result)
}

func TestWhenLengthAreTheSameAndHasOccurrenceThenReturnOne(t *testing.T) {
	// Arrange
	haystack := "abc"
	needle := "abc"

	// Act
	result := strStr(haystack, needle)

	// Assert
	expectedResult := 0
	assert.Equal(t, expectedResult, result)
}

func TestWhenLengthAreTheSameAndHasNoOccurrenceThenReturnOne(t *testing.T) {
	// Arrange
	haystack := "abc"
	needle := "def"

	// Act
	result := strStr(haystack, needle)

	// Assert
	expectedResult := -1
	assert.Equal(t, expectedResult, result)
}

func TestWhenOneCharAndHasOccurrenceThenReturnOne(t *testing.T) {
	// Arrange
	haystack := "a"
	needle := "a"

	// Act
	result := strStr(haystack, needle)

	// Assert
	expectedResult := 0
	assert.Equal(t, expectedResult, result)
}

func TestWhenOneCharAndHasNoOccurrenceThenReturnOne(t *testing.T) {
	// Arrange
	haystack := "a"
	needle := "b"

	// Act
	result := strStr(haystack, needle)

	// Assert
	expectedResult := -1
	assert.Equal(t, expectedResult, result)
}

func TestWhen(t *testing.T) {
	// Arrange
	haystack := "mississippi"
	needle := "issip"

	// Act
	result := strStr(haystack, needle)

	// Assert
	expectedResult := 4
	assert.Equal(t, expectedResult, result)
}
