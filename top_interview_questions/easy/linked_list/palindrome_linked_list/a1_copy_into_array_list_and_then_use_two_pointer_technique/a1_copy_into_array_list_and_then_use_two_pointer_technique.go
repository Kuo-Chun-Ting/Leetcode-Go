package a1_copy_into_array_list_and_then_use_two_pointer_technique

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	arr := listNodeToArray(head)

	start := 0
	end := len(arr) - 1
	for start < end {
		if arr[start] != arr[end] {
			return false
		}
		start++
		end--
	}

	return true
}

func listNodeToArray(head *ListNode) []int {
	arr := []int{}
	currNode := head
	for currNode != nil {
		arr = append(arr, currNode.Val)
		currNode = currNode.Next
	}
	return arr
}
