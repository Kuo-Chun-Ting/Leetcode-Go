package a2_space_optimal_operation_sub_optimal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhen01031Return131200(t *testing.T) {
	// Arrange
	nums := []int{0, 1, 0, 3, 12}

	// Act
	moveZeroes(nums)

	// Assert
	assert.Equal(t, []int{1, 3, 12, 0, 0}, nums)
}

func TestWhen0Return0(t *testing.T) {
	// Arrange
	nums := []int{0}

	// Act
	moveZeroes(nums)

	// Assert
	assert.Equal(t, []int{0}, nums)
}
