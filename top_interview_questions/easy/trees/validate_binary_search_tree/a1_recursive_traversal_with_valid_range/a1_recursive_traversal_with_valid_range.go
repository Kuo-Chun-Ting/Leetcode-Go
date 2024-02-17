package a1_recursive_traversal_with_valid_range

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validateBST(root, math.MinInt, math.MaxInt)
}

func validateBST(root *TreeNode, low, high int) bool {
	if root == nil {
		return true
	}

	if (root.Val <= low) || (root.Val >= high) {
		return false
	}

	return validateBST(root.Left, low, root.Val) && validateBST(root.Right, root.Val, high)
}
