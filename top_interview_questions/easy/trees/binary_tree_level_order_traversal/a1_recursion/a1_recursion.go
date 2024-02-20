package a1_recursion

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	levelOrder := [][]int{}
	currLevel := 0
	preorderDFS(root, currLevel, &levelOrder)
	return levelOrder
}

func preorderDFS(root *TreeNode, currLevel int, levelOrder *[][]int) {
	if root == nil {
		return
	}

	if len(*levelOrder) == currLevel {
		*levelOrder = append(*levelOrder, []int{})
	}

	(*levelOrder)[currLevel] = append((*levelOrder)[currLevel], root.Val)
	currLevel += 1

	preorderDFS(root.Left, currLevel, levelOrder)
	preorderDFS(root.Right, currLevel, levelOrder)
}
