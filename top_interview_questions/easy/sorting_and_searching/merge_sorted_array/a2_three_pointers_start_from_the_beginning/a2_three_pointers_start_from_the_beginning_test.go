package a2_three_pointers_start_from_the_beginning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenAllNums1SmallerThanNums2(t *testing.T) {
	// Arrange
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{4, 5, 6}

	// Act
	merge(nums1, 3, nums2, 3)

	// Assert
	expectedResult := []int{1, 2, 3, 4, 5, 6}
	assert.Equal(t, expectedResult, nums1)
}

func TestWhenAllNums1GreaterThanNums2(t *testing.T) {
	// Arrange
	nums1 := []int{4, 5, 6, 0, 0, 0}
	nums2 := []int{1, 2, 3}

	// Act
	merge(nums1, 3, nums2, 3)

	// Assert
	expectedResult := []int{1, 2, 3, 4, 5, 6}
	assert.Equal(t, expectedResult, nums1)
}
