package a1_recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenRootNil(t *testing.T) {
	// Act
	levelOrder := levelOrder(nil)

	// Assert
	expectedLevelOrder := [][]int{}
	assert.Equal(t, expectedLevelOrder, levelOrder)
}

func TestWhenOnlyRoot(t *testing.T) {
	// Arrange
	//      1
	root := &TreeNode{Val: 1}

	// Act
	levelOrder := levelOrder(root)

	// Assert
	expectedLevelOrder := [][]int{
		{1},
	}
	assert.Equal(t, expectedLevelOrder, levelOrder)
}

func TestWhenTwoLevelTree(t *testing.T) {
	// Arrange
	//      2
	//    1   3
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	// Act
	levelOrder := levelOrder(root)

	// Assert
	expectedLevelOrder := [][]int{
		{2}, {1, 3},
	}
	assert.Equal(t, expectedLevelOrder, levelOrder)
}

func TestWhenThreeLevelTree(t *testing.T) {
	// Arrange
	//      5
	//    3   7
	//   1 4 6 9
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}

	// Act
	levelOrder := levelOrder(root)

	// Assert
	expectedLevelOrder := [][]int{
		{5}, {3, 7}, {1, 4, 6, 9},
	}
	assert.Equal(t, expectedLevelOrder, levelOrder)
}

func TestWhenSkewedTree(t *testing.T) {
	// Arrange
	//      5
	//    3
	//   1
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
		},
	}

	// Act
	levelOrder := levelOrder(root)

	// Assert
	expectedLevelOrder := [][]int{
		{5}, {3}, {1},
	}
	assert.Equal(t, expectedLevelOrder, levelOrder)
}
