package a1_preorder_traversal_always_choose_middle_left_node_as_root

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	var helper func(left, right int) *TreeNode
	helper = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		rootIdx := (left + right) / 2
		root := TreeNode{Val: nums[rootIdx]}
		root.Left = helper(left, rootIdx-1)
		root.Right = helper(rootIdx+1, right)
		return &root
	}

	return helper(0, len(nums)-1)
}
