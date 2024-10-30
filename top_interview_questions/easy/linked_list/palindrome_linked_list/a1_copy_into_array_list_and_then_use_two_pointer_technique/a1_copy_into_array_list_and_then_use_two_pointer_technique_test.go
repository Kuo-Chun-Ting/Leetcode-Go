package a1_copy_into_array_list_and_then_use_two_pointer_technique

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenOneNode(t *testing.T) {
	// Arrange
	node := &ListNode{Val: 1}

	// Act
	result := isPalindrome(node)

	// Assert
	assert.Equal(t, true, result)
}

func TestWhenTwoNodeAndNotPalindrome(t *testing.T) {
	// Arrange
	node := &ListNode{
		Val: 1, Next: &ListNode{
			Val: 2,
		},
	}

	// Act
	result := isPalindrome(node)

	// Assert
	assert.Equal(t, false, result)
}

func TestWhenTwoNodeAndIsPalindrome(t *testing.T) {
	// Arrange
	node := &ListNode{
		Val: 1, Next: &ListNode{
			Val: 1,
		},
	}

	// Act
	result := isPalindrome(node)

	// Assert
	assert.Equal(t, true, result)
}

func TestWhenThreeNodeAnNotsPalindrome(t *testing.T) {
	// Arrange
	node := &ListNode{
		Val: 1, Next: &ListNode{
			Val: 2, Next: &ListNode{
				Val: 3,
			},
		},
	}

	// Act
	result := isPalindrome(node)

	// Assert
	assert.Equal(t, false, result)
}

func TestWhenThreeNodeAndIsPalindrome(t *testing.T) {
	// Arrange
	node := &ListNode{
		Val: 1, Next: &ListNode{
			Val: 2, Next: &ListNode{
				Val: 1,
			},
		},
	}

	// Act
	result := isPalindrome(node)

	// Assert
	assert.Equal(t, true, result)
}
