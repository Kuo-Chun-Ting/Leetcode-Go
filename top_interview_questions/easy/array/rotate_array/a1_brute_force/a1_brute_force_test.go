package a1_brute_force

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateWhenOneItems(t *testing.T) {
	// Arrange
	nums := []int{1}
	k := 3

	// Act
	rotate(nums, k)

	// Assert
	assert.Equal(t, nums, nums)
}
