package a1_rotate_groups_of_four_cells

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenNEqual3(t *testing.T) {
	// Arrange
	matrix := [][]int{
		{5, 1, 9},
		{2, 4, 8},
		{13, 3, 6},
	}

	// Act
	rotate(matrix)

	// Assert
	expectedMatrix := [][]int{
		{13, 2, 5},
		{3, 4, 1},
		{6, 8, 9},
	}
	assert.Equal(t, expectedMatrix, matrix)
}

func TestWhenNEqual4(t *testing.T) {
	// Arrange
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	// Act
	rotate(matrix)

	// Assert
	expectedMatrix := [][]int{
		{15, 13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7, 10, 11},
	}
	assert.Equal(t, expectedMatrix, matrix)
}

func TestWhenNEqual5(t *testing.T) {
	// Arrange
	matrix := [][]int{
		{15, 13, 2, 5, 2},
		{14, 3, 4, 1, 2},
		{12, 6, 8, 9, 2},
		{16, 7, 10, 11, 2},
		{19, 17, 12, 1, 2},
	}

	// Act
	rotate(matrix)

	// Assert
	expectedMatrix := [][]int{
		{19, 16, 12, 14, 15},
		{17, 7, 6, 3, 13},
		{12, 10, 8, 4, 2},
		{1, 11, 9, 1, 5},
		{2, 2, 2, 2, 2},
	}
	assert.Equal(t, expectedMatrix, matrix)
}
