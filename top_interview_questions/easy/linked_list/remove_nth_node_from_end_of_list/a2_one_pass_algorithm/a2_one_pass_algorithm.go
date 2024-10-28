package a1_two_pass_algorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	firstPointer := &ListNode{Val: -1, Next: head}
	seconedPointer := &ListNode{Val: -1, Next: head}

	count := 0
	for count < n {
		firstPointer = firstPointer.Next
		count++
	}

	for firstPointer.Next != nil {
		firstPointer = firstPointer.Next
		seconedPointer = seconedPointer.Next
	}

	if seconedPointer.Val == -1 {
		head = head.Next
	} else {
		seconedPointer.Next = seconedPointer.Next.Next
	}

	return head
}
