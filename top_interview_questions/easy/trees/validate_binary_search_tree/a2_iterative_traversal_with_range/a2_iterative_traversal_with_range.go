package a1_recursive_traversal_with_valid_range

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type StackItem struct {
	Node *TreeNode
	Low  int
	High int
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	validatedRoot := StackItem{
		Node: root,
		Low:  math.MinInt,
		High: math.MaxInt,
	}
	stack := []StackItem{validatedRoot}

	for {
		if len(stack) == 0 {
			break
		}

		validatedNode := stack[0]
		stack = stack[1:]

		if validatedNode.Node.Val <= validatedNode.Low || validatedNode.Node.Val >= validatedNode.High {
			return false
		}

		if validatedNode.Node.Left != nil {
			leftToValidate := StackItem{
				Node: validatedNode.Node.Left,
				Low:  validatedNode.Low,
				High: validatedNode.Node.Val,
			}
			stack = append(stack, leftToValidate)
		}

		if validatedNode.Node.Right != nil {
			rightToValidate := StackItem{
				Node: validatedNode.Node.Right,
				Low:  validatedNode.Node.Val,
				High: validatedNode.High,
			}
			stack = append(stack, rightToValidate)
		}
	}

	return true
}
