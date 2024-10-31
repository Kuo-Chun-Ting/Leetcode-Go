package a1_recursive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenOneLevel(t *testing.T) {
	// Arrange
	root := &TreeNode{Val: 1}

	// Act
	result := isSymmetric(root)

	// Assert
	assert.Equal(t, true, result)
}

func TestWhenTwoLevelsAndNotSymmetricThenReturnFalse(t *testing.T) {
	// Arrange
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 1,
		},
	}

	// Act
	result := isSymmetric(root)

	// Assert
	assert.Equal(t, false, result)
}

func TestWhenTwoLevelsAndIsSymmetricThenReturnTrue(t *testing.T) {
	// Arrange
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 1,
		},
	}

	// Act
	result := isSymmetric(root)

	// Assert
	assert.Equal(t, true, result)
}

func TestWhenThreeLevelsAndNotSymmetricThenReturnFalse(t *testing.T) {
	// Arrange
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}

	// Act
	result := isSymmetric(root)

	// Assert
	assert.Equal(t, false, result)
}

func TestWhenThreeLevelsAndIsSymmetricThenReturnTrue(t *testing.T) {
	// Arrange
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}

	// Act
	result := isSymmetric(root)

	// Assert
	assert.Equal(t, true, result)
}
