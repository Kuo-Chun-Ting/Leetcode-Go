package a2_tail_recursion_dfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	currDepth := 0
	var maxDepth *int = &currDepth

	maxDepthHelper(root, currDepth, maxDepth)
	return *maxDepth
}

func maxDepthHelper(root *TreeNode, currentDepth int, maxHeight *int) int {
	if root == nil {
		if currentDepth > *maxHeight {
			*maxHeight = currentDepth
		}

		return currentDepth
	}

	leftHeight := maxDepthHelper(root.Left, currentDepth+1, maxHeight)
	rightHeight := maxDepthHelper(root.Right, currentDepth+1, maxHeight)

	return max(leftHeight, rightHeight)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
