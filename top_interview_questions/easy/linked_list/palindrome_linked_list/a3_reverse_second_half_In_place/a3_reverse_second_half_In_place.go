package a1_copy_into_array_list_and_then_use_two_pointer_technique

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}

	fastPtr := &ListNode{Val: -1, Next: head}
	slowPtr := &ListNode{Val: -1, Next: head}
	palindromeLength := 0

	for fastPtr != nil && fastPtr.Next != nil {
		fastPtr = fastPtr.Next.Next
		slowPtr = slowPtr.Next
		palindromeLength++
	}

	tail := slowPtr.Next
	reversedTail := reverseListNode(tail)

	for reversedTail != nil {
		if head.Val != reversedTail.Val {
			return false
		}
		head = head.Next
		reversedTail = reversedTail.Next
	}
	return true
}

func reverseListNode(head *ListNode) *ListNode {
	var prev *ListNode = nil
	var next *ListNode = head.Next

	for next != nil {
		head.Next = prev
		prev = head
		head = next
		next = head.Next
	}
	head.Next = prev
	return head
}
