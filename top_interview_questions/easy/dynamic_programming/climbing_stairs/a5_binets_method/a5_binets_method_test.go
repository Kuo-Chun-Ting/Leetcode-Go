package a5_binets_method

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenOneStair(t *testing.T) {
	// Arrange
	n := 1
	// Act
	combination := climbStairs(n)

	// Assert
	assert.Equal(t, 1, combination)
}

func TestWhenTwoStairs(t *testing.T) {
	// Arrange
	n := 2
	// Act
	combination := climbStairs(n)

	// Assert
	assert.Equal(t, 2, combination)
}

func TestWhenThreeStairs(t *testing.T) {
	// Arrange
	n := 3
	// Act
	combination := climbStairs(n)

	// Assert
	assert.Equal(t, 3, combination)
}
