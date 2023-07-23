package reverse_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteNodeWhenNodeIsHead(t *testing.T) {
	// Arrange
	head := ListNode{
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
	deleteNode(&head)

	// Assert
	expectedHead := ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  3,
			Next: nil,
		},
	}
	assert.Equal(t, expectedHead, head)
}

func TestDeleteNodeWhenNodeInTheMiddle(t *testing.T) {
	// Arrange
	head := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	nodeToBeDeleted := head.Next

	// Act
	deleteNode(nodeToBeDeleted)

	// Assert
	expectedHead := ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  3,
			Next: nil,
		},
	}
	assert.Equal(t, expectedHead, head)
}
