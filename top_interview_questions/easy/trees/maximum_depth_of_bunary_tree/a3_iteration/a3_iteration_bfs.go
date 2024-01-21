package a3_iteration

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeHeightPair struct {
	Node   *TreeNode
	Height int
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	rootPair := &NodeHeightPair{Node: root, Height: 1}
	queue := []*NodeHeightPair{}
	maxHeight := 0

	// enqueue
	queue = append(queue, rootPair)

	for {
		if len(queue) == 0 {
			break
		}

		// dequeue
		pair := queue[0]
		queue = queue[1:]

		currNode := pair.Node
		currHeight := pair.Height

		if currHeight > maxHeight {
			maxHeight = currHeight
		}

		if currNode.Left != nil {
			leftNodePair := createNodeHeightPair(currNode.Left, currHeight+1)
			queue = append(queue, leftNodePair)
		}

		if currNode.Right != nil {
			rightNodePair := createNodeHeightPair(currNode.Right, currHeight+1)
			queue = append(queue, rightNodePair)
		}
	}

	return maxHeight
}

func createNodeHeightPair(node *TreeNode, height int) *NodeHeightPair {
	return &NodeHeightPair{
		Node:   node,
		Height: height,
	}
}
