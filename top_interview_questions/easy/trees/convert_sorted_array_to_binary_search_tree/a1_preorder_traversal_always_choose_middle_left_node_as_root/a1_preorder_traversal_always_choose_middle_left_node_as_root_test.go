package a1_preorder_traversal_always_choose_middle_left_node_as_root

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenOnlyRoot(t *testing.T) {
	// Arrange
	sortedArr := []int{1}

	// Act
	bst := sortedArrayToBST(sortedArr)

	// Assert
	expectedBST := &TreeNode{Val: 1}
	assert.Equal(t, expectedBST, bst)
}

func TestWhenOddNums(t *testing.T) {
	// Arrange
	sortedArr := []int{1, 2, 3}

	// Act
	bst := sortedArrayToBST(sortedArr)

	// Assert
	expectedBST := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}
	assert.Equal(t, expectedBST, bst)
}

func TestWhenEvenNums(t *testing.T) {
	// Arrange
	sortedArr := []int{1, 2}

	// Act
	bst := sortedArrayToBST(sortedArr)

	// Assert
	expectedBST := &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 2},
	}
	assert.Equal(t, expectedBST, bst)
}
