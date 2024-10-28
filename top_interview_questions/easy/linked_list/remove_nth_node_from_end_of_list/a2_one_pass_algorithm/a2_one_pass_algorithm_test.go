package a1_two_pass_algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenSizeOneReturnNil(t *testing.T) {
	// Arrange
	head := &ListNode{Val: 1}
	n := 1

	// Act
	result := removeNthFromEnd(head, n)

	// Assert
	assert.Nil(t, result)
}

func TestWhenRemoveFirstNodeThenResultShouldBeCorrect(t *testing.T) {
	// Arrange
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	n := 3

	// Act
	result := removeNthFromEnd(head, n)

	// Assert
	expectedResult := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
		},
	}
	assert.Equal(t, expectedResult, result)
}

func TestWhenRemoveLastNodeThenResultShouldBeCorrect(t *testing.T) {
	// Arrange
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	n := 1

	// Act
	result := removeNthFromEnd(head, n)

	// Assert
	expectedResult := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	assert.Equal(t, expectedResult, result)
}

func TestRemoveNthFromEndWhenRemoveMiddleNodeThenResultShouldBeCorrect(t *testing.T) {
	// Arrange
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	n := 2

	// Act
	result := removeNthFromEnd(head, n)

	// Assert
	expectedResult := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
		},
	}
	assert.Equal(t, expectedResult, result)
}
