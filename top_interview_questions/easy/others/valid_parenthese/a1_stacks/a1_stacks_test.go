package a1_stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenSinglePairParenthesesThenTrue(t *testing.T) {
	// Arrange
	input := "()"

	// Act
	result := isValid(input)

	// Assert
	assert.Equal(t, true, result)
}

func TestWhenMultipleDifferentParenthesesThenTrue(t *testing.T) {
	// Arrange
	input := "()[]{}"

	// Act
	result := isValid(input)

	// Assert
	assert.Equal(t, true, result)
}

func TestWhenMismatchedParenthesesThenFalse(t *testing.T) {
	// Arrange
	input := "()[{}"

	// Act
	result := isValid(input)

	// Assert
	assert.Equal(t, false, result)
}

func TestWhenNestedParenthesesThenTrue(t *testing.T) {
	// Arrange
	input := "([{}])"

	// Act
	result := isValid(input)

	// Assert
	assert.Equal(t, true, result)
}

func TestWhenInputHasOnlyClosingBracketThenReturnFalse(t *testing.T) {
	// Arrange
	input := "]"

	// Act
	result := isValid(input)

	// Assert
	assert.Equal(t, false, result)
}
