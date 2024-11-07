package a1_follow_the_rules

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenOnlyPositiveNumbers(t *testing.T) {
	// Arrange
	str := "42"

	// Act
	result := myAtoi(str)

	// Assert
	assert.Equal(t, 42, result)
}

func TestWhenOnlyNegativeNumbers(t *testing.T) {
	// Arrange
	str := "-42"

	// Act
	result := myAtoi(str)

	// Assert
	assert.Equal(t, -42, result)
}

func TestWhenPositiveNumbersContainingWhiteSpace(t *testing.T) {
	// Arrange
	str := " 42 "

	// Act
	result := myAtoi(str)

	// Assert
	assert.Equal(t, 42, result)
}

func TestWhenOnlyNegativeNumbersContainingWhiteSpace(t *testing.T) {
	// Arrange
	str := " -42 "

	// Act
	result := myAtoi(str)

	// Assert
	assert.Equal(t, -42, result)
}

func TestWhenPosiveNumberAndPlusInMiddleThenSkipEndOfNumber(t *testing.T) {
	// Arrange
	str := "0+1"

	// Act
	result := myAtoi(str)

	// Assert
	assert.Equal(t, 0, result)
}

func TestWhenNegativeNumbersAndMinusInMiddleThenSkipEndOfNumber(t *testing.T) {
	// Arrange
	str := "0-1"

	// Act
	result := myAtoi(str)

	// Assert
	assert.Equal(t, 0, result)
}

func TestWhenStartWithNonNumberThenReturn0(t *testing.T) {
	// Arrange
	str := "a123"

	// Act
	result := myAtoi(str)

	// Assert
	assert.Equal(t, 0, result)
}

func TestWhenEndWithNonNumberThenReturn0(t *testing.T) {
	// Arrange
	str := " 123abc"

	// Act
	result := myAtoi(str)

	// Assert
	assert.Equal(t, 123, result)
}

func TestWhenResultLessThanMinThenReturnMin(t *testing.T) {
	// Arrange
	str := " -2147483649"

	// Act
	result := myAtoi(str)

	// Assert
	assert.Equal(t, math.MinInt32, result)
}

func TestWhenResultGreaterThanMaxThenReturnMax(t *testing.T) {
	// Arrange
	str := " 2147483648"

	// Act
	result := myAtoi(str)

	// Assert
	assert.Equal(t, math.MaxInt32, result)
}

func TestWhenConsequenceZeroCauseOverfloatThenReturnBoundary(t *testing.T) {
	// Arrange
	str := "10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459"

	// Act
	result := myAtoi(str)

	// Assert
	assert.Equal(t, math.MaxInt32, result)
}

func TestWhenStartWithZeroThenSkipStartingZero(t *testing.T) {
	// Arrange
	str := "  0000000000012345678"

	// Act
	result := myAtoi(str)

	// Assert
	assert.Equal(t, 12345678, result)
}

func TestWhenMax(t *testing.T) {
	// Arrange
	str := "  2147483647"

	// Act
	result := myAtoi(str)

	// Assert
	assert.Equal(t, 2147483647, result)
}

func TestWhenMin(t *testing.T) {
	// Arrange
	str := "  -2147483647"

	// Act
	result := myAtoi(str)

	// Assert
	assert.Equal(t, -2147483647, result)
}
