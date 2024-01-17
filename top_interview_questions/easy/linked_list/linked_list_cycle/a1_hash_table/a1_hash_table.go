package a1_hash_table

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	seen := make(map[*ListNode]bool)

	for {
		if head == nil {
			break
		}

		if _, ok := seen[head]; ok {
			return true
		}

		seen[head] = true
		head = head.Next
	}

	return false
}
