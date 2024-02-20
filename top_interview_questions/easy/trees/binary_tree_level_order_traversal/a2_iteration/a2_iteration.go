package a2_iteration

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type QueueItem struct {
	Node  *TreeNode
	Level int
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	levelOrder := [][]int{}
	queue := []QueueItem{}
	rootItem := QueueItem{
		Node:  root,
		Level: 0,
	}
	queue = append(queue, rootItem)

	for {
		if len(queue) == 0 {
			break
		}

		item := queue[0]
		queue = queue[1:]
		if len(levelOrder) == item.Level {
			levelOrder = append(levelOrder, []int{})
		}
		levelOrder[item.Level] = append(levelOrder[item.Level], item.Node.Val)

		if item.Node.Left != nil {
			leftItem := QueueItem{
				Node:  item.Node.Left,
				Level: item.Level + 1,
			}
			queue = append(queue, leftItem)
		}

		if item.Node.Right != nil {
			rightItem := QueueItem{
				Node:  item.Node.Right,
				Level: item.Level + 1,
			}
			queue = append(queue, rightItem)
		}
	}

	return levelOrder
}
