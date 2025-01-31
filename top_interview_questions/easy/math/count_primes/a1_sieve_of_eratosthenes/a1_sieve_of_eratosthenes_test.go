package a1_sieve_of_eratosthenes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhen0ThenReturn0(t *testing.T) {
	// Arrange
	n := 0

	// Act
	result := countPrimes(n)

	// Assert
	expectedResult := 0
	assert.Equal(t, expectedResult, result)
}

func TestWhen1ThenReturn0(t *testing.T) {
	// Arrange
	n := 1

	// Act
	result := countPrimes(n)

	// Assert
	expectedResult := 0
	assert.Equal(t, expectedResult, result)
}

func TestWhen2ThenReturn1(t *testing.T) {
	// Arrange
	n := 2

	// Act
	result := countPrimes(n)

	// Assert
	expectedResult := 0
	assert.Equal(t, expectedResult, result)
}

func TestWhen3ThenReturn1(t *testing.T) {
	// Arrange
	n := 3

	// Act
	result := countPrimes(n)

	// Assert
	expectedResult := 1
	assert.Equal(t, expectedResult, result)
}

func TestWhen4ThenReturn1(t *testing.T) {
	// Arrange
	n := 4

	// Act
	result := countPrimes(n)

	// Assert
	expectedResult := 2
	assert.Equal(t, expectedResult, result)
}

func TestWhen10ThenReturn4(t *testing.T) {
	// Arrange
	n := 10

	// Act
	result := countPrimes(n)

	// Assert
	expectedResult := 4
	assert.Equal(t, expectedResult, result)
}
