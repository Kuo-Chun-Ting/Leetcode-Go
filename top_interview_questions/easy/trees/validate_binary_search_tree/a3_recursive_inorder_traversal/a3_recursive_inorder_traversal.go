package a3_recursive_inorder_traversal

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	lastVal := math.MinInt
	return inorder(root, &lastVal)
}

func inorder(root *TreeNode, lastVal *int) bool {
	if root == nil {
		return true
	}

	if !inorder(root.Left, lastVal) {
		return false
	}

	if root.Val <= *lastVal {
		return false
	} else {
		*lastVal = root.Val
	}

	return inorder(root.Right, lastVal)
}
