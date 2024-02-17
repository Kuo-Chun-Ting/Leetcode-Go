package a4_iterative_inorder_traversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenRootNil(t *testing.T) {
	// Arrange
	var root *TreeNode

	// Act
	hasCycle := isValidBST(root)

	// Assert
	assert.Equal(t, true, hasCycle)
}

func TestWhenOnlyRoot(t *testing.T) {
	// Arrange
	root := &TreeNode{Val: 1}

	// Act
	hasCycle := isValidBST(root)

	// Assert
	assert.Equal(t, true, hasCycle)
}

func TestWhenTwoLevelsAndNotBST(t *testing.T) {
	// Arrange
	//      1
	//    2   3
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}

	// Act
	hasCycle := isValidBST(root)

	// Assert
	assert.Equal(t, false, hasCycle)
}

func TestWhenTwoLevelsAndIsBST(t *testing.T) {
	// Arrange
	//      2
	//    1   3
	root := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}

	// Act
	hasCycle := isValidBST(root)

	// Assert
	assert.Equal(t, true, hasCycle)
}

func TestWhenTwoLevelsAndAllValuesAreTheSame(t *testing.T) {
	// Arrange
	//      2
	//    2   2
	root := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 2},
	}

	// Act
	hasCycle := isValidBST(root)

	// Assert
	assert.Equal(t, false, hasCycle)
}

func TestWhenTwoLevelsAndNotCompleteBST(t *testing.T) {
	// Arrange
	//      0
	//    -1
	root := &TreeNode{
		Val:  0,
		Left: &TreeNode{Val: -1},
	}

	// Act
	hasCycle := isValidBST(root)

	// Assert
	assert.Equal(t, true, hasCycle)
}

func TestWhenThreeLevelsAndNotBST(t *testing.T) {
	// Arrange
	//      5
	//    3   7
	//   2 1 6 8
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}

	// Act
	hasCycle := isValidBST(root)

	// Assert
	assert.Equal(t, false, hasCycle)
}

func TestWhenThreeLevelsAndIsBST(t *testing.T) {
	// Arrange
	//      5
	//    3   7
	//   2 4 6 8
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
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
				Val: 8,
			},
		},
	}

	// Act
	hasCycle := isValidBST(root)

	// Assert
	assert.Equal(t, true, hasCycle)
}

func TestWhenThreeLevelsAndAllValuesAreTheSame(t *testing.T) {
	// Arrange
	//      2
	//    2   2
	//   2 2 2 2
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
	}

	// Act
	hasCycle := isValidBST(root)

	// Assert
	assert.Equal(t, false, hasCycle)
}

func TestWhenThreeLevelsAndNotCompleteBST(t *testing.T) {
	// Arrange
	//      5
	//    2
	//   1 3
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	// Act
	hasCycle := isValidBST(root)

	// Assert
	assert.Equal(t, true, hasCycle)
}

func TestWhenLeftTreesRightNodeGreaterThanHigherParent(t *testing.T) {
	// Arrange
	//      5
	//    4   6
	//   3 7
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 6,
		},
	}

	// Act
	hasCycle := isValidBST(root)

	// Assert
	assert.Equal(t, false, hasCycle)
}

func TestWhenRightTreesLeftNodeLessThanHigherParent(t *testing.T) {
	// Arrange
	//      5
	//    4   6
	//   3 7 2
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 2,
			},
		},
	}

	// Act
	hasCycle := isValidBST(root)

	// Assert
	assert.Equal(t, false, hasCycle)
}
