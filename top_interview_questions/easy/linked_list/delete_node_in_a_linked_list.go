package reverse_string

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	next := node.Next
	node.Val = next.Val
	node.Next = next.Next
}
