package a1_two_pass_algorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	headLength := countLength(head)
	indexToRemoveFromStart := headLength - n
	currIndex := 0
	currNode := head

	for currNode != nil {
		if indexToRemoveFromStart == 0 {
			return head.Next
		} else if currIndex == indexToRemoveFromStart-1 {
			nextNode := currNode.Next.Next
			currNode.Next = currNode.Next.Next
			currNode = nextNode
		} else {
			currNode = currNode.Next
		}

		currIndex++
	}

	return head
}

func countLength(node *ListNode) int {
	length := 0
	currNode := node
	for currNode != nil {
		currNode = currNode.Next
		length++
	}

	return length
}
