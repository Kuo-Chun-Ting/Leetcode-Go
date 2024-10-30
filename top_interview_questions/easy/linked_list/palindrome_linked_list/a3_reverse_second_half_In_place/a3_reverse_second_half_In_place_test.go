package a1_copy_into_array_list_and_then_use_two_pointer_technique

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseListNodeWhenOneNode(t *testing.T) {
	// Arrange
	node := &ListNode{Val: 1}

	// Act
	result := reverseListNode(node)

	// Assert
	expectedResult := &ListNode{Val: 1}
	assert.Equal(t, expectedResult, result)
}

func TestReverseListNodeWhenTwoNode(t *testing.T) {
	// Arrange
	node := &ListNode{
		Val: 1, Next: &ListNode{
			Val: 2,
		},
	}

	// Act
	result := reverseListNode(node)

	// Assert
	expectedResult := &ListNode{
		Val: 2, Next: &ListNode{
			Val: 1,
		},
	}
	assert.Equal(t, expectedResult, result)
}

func TestReverseListNodeWhenThreeNode(t *testing.T) {
	// Arrange
	node := &ListNode{
		Val: 1, Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	// Act
	result := reverseListNode(node)

	// Assert
	expectedResult := &ListNode{
		Val: 3, Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
			},
		},
	}
	assert.Equal(t, expectedResult, result)
}

func TestWhenOneNode(t *testing.T) {
	// Arrange
	node := &ListNode{Val: 1}

	// Act
	result := isPalindrome(node)

	// Assert
	assert.Equal(t, true, result)
}

func TestWhenEvenNodeAndIsPalindromeThenReturnTrue(t *testing.T) {
	// Arrange
	node := &ListNode{
		Val: 1, Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1, Next: &ListNode{
					Val: 1,
				},
			},
		},
	}

	// Act
	result := isPalindrome(node)

	// Assert
	assert.Equal(t, true, result)
}

func TestWhenEvenNodeAndNotPalindromeThenReturnFalse(t *testing.T) {
	// Arrange
	node := &ListNode{
		Val: 1, Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2, Next: &ListNode{
					Val: 1,
				},
			},
		},
	}

	// Act
	result := isPalindrome(node)

	// Assert
	assert.Equal(t, false, result)
}

func TestWhenOddNodeAndIsPalindromeThenReturnTrue(t *testing.T) {
	// Arrange
	node := &ListNode{
		Val: 1, Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2, Next: &ListNode{
						Val: 1,
					},
				},
			},
		},
	}

	// Act
	result := isPalindrome(node)

	// Assert
	assert.Equal(t, true, result)
}

func TestWhenOddNodeAndNotPalindromeThenReturnFalse(t *testing.T) {
	// Arrange
	node := &ListNode{
		Val: 1, Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 3, Next: &ListNode{
						Val: 1,
					},
				},
			},
		},
	}

	// Act
	result := isPalindrome(node)

	// Assert
	assert.Equal(t, false, result)
}
