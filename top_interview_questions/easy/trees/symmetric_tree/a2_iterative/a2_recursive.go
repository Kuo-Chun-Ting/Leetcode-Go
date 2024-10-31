package a2_iterative

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	queue := []*TreeNode{root, root}

	for len(queue) > 0 {
		t1 := queue[0]
		t2 := queue[1]
		queue = queue[2:]

		if t1 == nil && t2 == nil {
			continue
		}

		if t1 == nil || t2 == nil {
			return false
		}

		if t1.Val != t2.Val {
			return false
		}

		queue = append(queue, t1.Left, t2.Right, t1.Right, t2.Left)
	}

	return true
}
