package a1_recursion

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	} else {
		return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
