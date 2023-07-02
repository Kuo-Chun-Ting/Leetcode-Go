package reverse_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseStringWhenStringIsNillThenDoNothing(t *testing.T) {
	// Arrange
	var str []byte

	// Act
	reverseString(str)

	// Assert
	var expectedStr []byte
	assert.Equal(t, expectedStr, str)
}

func TestReverseStringWhenOnlyOneCharacterThenResultShouldBeCorrect(t *testing.T) {
	// Arrange
	str := []byte{'h'}

	// Act
	reverseString(str)

	// Assert
	expectedStr := []byte{'h'}
	assert.Equal(t, expectedStr, str)
}

func TestReverseStringWhenOnlyTwoCharactersThenResultShouldBeCorrect(t *testing.T) {
	// Arrange
	str := []byte{'h', 'e'}

	// Act
	reverseString(str)

	// Assert
	expectedStr := []byte{'e', 'h'}
	assert.Equal(t, expectedStr, str)
}

func TestReverseStringWhenStringCountIsOddNumberThenResultShouldBeCorrect(t *testing.T) {
	// Arrange
	str := []byte{'h', 'e', 'l', 'l', 'o'}

	// Act
	reverseString(str)

	// Assert
	expectedStr := []byte{'o', 'l', 'l', 'e', 'h'}
	assert.Equal(t, expectedStr, str)
}

func TestReverseStringWhenStringCountIsEvenNumberThenResultShouldBeCorrect(t *testing.T) {
	// Arrange
	str := []byte{'h', 'e', 'e', 'l', 'l', 'o'}

	// Act
	reverseString(str)

	// Assert
	expectedStr := []byte{'o', 'l', 'l', 'e', 'e', 'h'}
	assert.Equal(t, expectedStr, str)
}
