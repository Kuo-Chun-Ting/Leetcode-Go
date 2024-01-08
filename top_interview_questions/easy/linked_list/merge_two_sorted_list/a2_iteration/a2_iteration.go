package a2_recursion

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var newHead *ListNode

	if list1.Val < list2.Val {
		newHead = list1
		list1 = list1.Next
	} else {
		newHead = list2
		list2 = list2.Next
	}

	sortingNode := newHead

	for {
		if list1 == nil {
			sortingNode.Next = list2
			break
		}

		if list2 == nil {
			sortingNode.Next = list1
			break
		}

		if list1.Val < list2.Val {
			sortingNode.Next = list1
			sortingNode = sortingNode.Next
			list1 = list1.Next
		} else {
			sortingNode.Next = list2
			sortingNode = sortingNode.Next
			list2 = list2.Next
		}
	}

	return newHead
}
