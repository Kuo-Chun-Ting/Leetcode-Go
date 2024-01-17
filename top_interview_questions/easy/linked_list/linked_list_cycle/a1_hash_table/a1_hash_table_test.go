package a1_hash_table

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenHeadIsNil(t *testing.T) {
	// Arrange
	var Node1 *ListNode

	// Act
	hasCycle := hasCycle(Node1)

	// Assert
	assert.Equal(t, false, hasCycle)
}

func TestWhenNodesHasCycle(t *testing.T) {
	// Arrange
	Node1 := &ListNode{
		Val: 1,
	}
	Node2 := &ListNode{
		Val: 2,
	}
	Node3 := &ListNode{
		Val: 3,
	}
	Node1.Next = Node2
	Node2.Next = Node3
	Node3.Next = Node2

	// Act
	hasCycle := hasCycle(Node1)

	// Assert
	assert.Equal(t, true, hasCycle)
}

func TestWhenNodesHasNoCycle(t *testing.T) {
	// Arrange
	Node1 := &ListNode{
		Val: 1,
	}
	Node2 := &ListNode{
		Val: 2,
	}
	Node3 := &ListNode{
		Val: 3,
	}
	Node1.Next = Node2
	Node2.Next = Node3
	Node3.Next = nil

	// Act
	hasCycle := hasCycle(Node1)

	// Assert
	assert.Equal(t, false, hasCycle)
}
