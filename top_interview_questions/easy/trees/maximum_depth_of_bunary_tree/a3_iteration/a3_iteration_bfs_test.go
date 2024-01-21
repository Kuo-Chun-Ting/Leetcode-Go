package a3_iteration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenNoNode(t *testing.T) {
	// Arrange
	var root *TreeNode

	// Act
	hasCycle := maxDepth(root)

	// Assert
	assert.Equal(t, 0, hasCycle)
}

func TestWhenOnlyRoot(t *testing.T) {
	// Arrange
	root := &TreeNode{}

	// Act
	hasCycle := maxDepth(root)

	// Assert
	assert.Equal(t, 1, hasCycle)
}

func TestWhenSkewedTree(t *testing.T) {
	// Arrange
	root := &TreeNode{
		Left: &TreeNode{Left: &TreeNode{}},
	}

	// Act
	hasCycle := maxDepth(root)

	// Assert
	assert.Equal(t, 3, hasCycle)
}

func TestWhenLeftHeightEqualtoRightHeight(t *testing.T) {
	// Arrange
	root := &TreeNode{
		Left:  &TreeNode{Left: &TreeNode{}},
		Right: &TreeNode{Right: &TreeNode{}},
	}

	// Act
	hasCycle := maxDepth(root)

	// Assert
	assert.Equal(t, 3, hasCycle)
}

func TestWhenLeftHeightNotEqualtoRightHeight(t *testing.T) {
	// Arrange
	root := &TreeNode{
		Left:  &TreeNode{Left: &TreeNode{Left: &TreeNode{}}},
		Right: &TreeNode{Right: &TreeNode{}},
	}

	// Act
	hasCycle := maxDepth(root)

	// Assert
	assert.Equal(t, 4, hasCycle)
}
