package a4_iterative_inorder_traversal

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	lastVal := math.MinInt
	stack := []*TreeNode{}
	currNode := root

	for {
		if currNode == nil && len(stack) == 0 {
			break
		}

		for {
			if currNode == nil {
				break
			}

			stack = append(stack, currNode)
			currNode = currNode.Left
		}

		currNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if currNode.Val <= lastVal {
			return false
		}

		lastVal = currNode.Val
		currNode = currNode.Right
	}

	return true
}
