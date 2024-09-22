package a3_mask_and_shift

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenFirstAndLastBitsAreZeroThenOutputShouldBeCorrect(t *testing.T) {
	// Arrange
	input := uint32(1073741824) // 01000000000000000000000000000000 (First and last bits are 0)

	// Act
	result := reverseBits(input)

	// Assert
	expectedOutput := uint32(2) // Expected: 00000000000000000000000000000010
	assert.Equal(t, expectedOutput, result)
}

func TestWhenFirstBitIsZeroAndLastBitIsOneThenOutputShouldBeCorrect(t *testing.T) {
	// Arrange
	input := uint32(1073741825) // 01000000000000000000000000000001 (First bit is 0, last bit is 1)

	// Act
	result := reverseBits(input)

	// Assert
	expectedOutput := uint32(2147483650) // Expected: 10000000000000000000000000000010
	assert.Equal(t, expectedOutput, result)
}

func TestWhenFirstBitIsOneAndLastBitIsZeroThenOutputShouldBeCorrect(t *testing.T) {
	// Arrange
	input := uint32(2147483648) // 10000000000000000000000000000000 (First bit is 1, last bit is 0)

	// Act
	result := reverseBits(input)

	// Assert
	expectedOutput := uint32(1) // Expected: 00000000000000000000000000000001
	assert.Equal(t, expectedOutput, result)
}

func TestWhenFirstAndLastBitsAreOneThenOutputShouldBeCorrect(t *testing.T) {
	// Arrange
	input := uint32(2147483649) // 10000000000000000000000000000001 (First and last bits are 1)

	// Act
	result := reverseBits(input)

	// Assert
	expectedOutput := uint32(2147483649) // Expected: 10000000000000000000000000000001
	assert.Equal(t, expectedOutput, result)
}
