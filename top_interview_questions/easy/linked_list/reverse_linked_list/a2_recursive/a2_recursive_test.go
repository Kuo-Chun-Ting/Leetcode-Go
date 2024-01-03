package a2_recursive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenListNodeNil(t *testing.T) {
	// Arrange
	var head *ListNode

	// Act
	reversedHead := reverseList(head)

	// Assert
	var expectedHead *ListNode
	assert.Equal(t, expectedHead, reversedHead)
}

func TestWhenNormalCase(t *testing.T) {
	// Arrange
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	// Act
	reversedHead := reverseList(head)

	// Assert
	expectedHead := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}
	assert.Equal(t, expectedHead, reversedHead)
}
