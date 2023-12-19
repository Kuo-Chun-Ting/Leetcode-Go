package a1_hash_map

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenNums1EqualsToNums2AndLengthIsOne(t *testing.T) {
	// Arrange
	nums1 := []int{1}
	nums2 := []int{1}

	// Act
	result := intersect(nums1, nums2)

	// Assert
	expectedIntersectionNums := []int{1}
	assert.Equal(t, expectedIntersectionNums, result)
}

func TestWhenNums1EqualsToNums2AndLengthNotOne(t *testing.T) {
	// Arrange
	nums1 := []int{1, 2}
	nums2 := []int{1, 2}

	// Act
	result := intersect(nums1, nums2)

	// Assert
	expectedIntersectionNums := []int{1, 2}
	assert.Equal(t, expectedIntersectionNums, result)
}

func TestWhenIntersectionNumCountGreaterThenOne(t *testing.T) {
	// Arrange
	nums1 := []int{1, 4, 4}
	nums2 := []int{1, 4, 4}

	// Act
	result := intersect(nums1, nums2)

	// Assert
	expectedIntersectionNums := []int{1, 4, 4}
	assert.Equal(t, expectedIntersectionNums, result)
}

func TestWhenNums1IncludsNums2(t *testing.T) {
	// Arrange
	nums1 := []int{1, 4, 4, 5, 6}
	nums2 := []int{1, 4, 4}

	// Act
	result := intersect(nums1, nums2)

	// Assert
	expectedIntersectionNums := []int{1, 4, 4}
	assert.Equal(t, expectedIntersectionNums, result)
}

func TestWhenNums1IntersectsWithNums2(t *testing.T) {
	// Arrange
	nums1 := []int{1, 2, 2, 4}
	nums2 := []int{1, 4}

	// Act
	result := intersect(nums1, nums2)

	// Assert
	expectedIntersectionNums := []int{1, 4}
	assert.Equal(t, expectedIntersectionNums, result)
}

func TestWhenNums1AndNums2AreDisjoint(t *testing.T) {
	// Arrange
	nums1 := []int{1, 2, 2, 4}
	nums2 := []int{8, 8, 9}

	// Act
	result := intersect(nums1, nums2)

	// Assert
	expectedIntersectionNums := []int{}
	assert.Equal(t, expectedIntersectionNums, result)
}
