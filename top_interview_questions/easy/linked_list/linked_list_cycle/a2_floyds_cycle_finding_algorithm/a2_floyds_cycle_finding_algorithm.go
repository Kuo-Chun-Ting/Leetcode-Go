package a1_hash_table

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fastRunner, slowRunner := head, head

	for {
		if fastRunner == nil || fastRunner.Next == nil {
			return false
		}

		slowRunner = slowRunner.Next
		fastRunner = fastRunner.Next.Next

		if fastRunner == slowRunner {
			return true
		}
	}
}
