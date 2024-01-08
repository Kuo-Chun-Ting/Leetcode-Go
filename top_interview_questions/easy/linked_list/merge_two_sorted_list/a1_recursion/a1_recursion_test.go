package a1_recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenBothTwoListsNil(t *testing.T) {
	// Arrange
	var left *ListNode
	var right *ListNode

	// Act
	sortedList := mergeTwoLists(left, right)

	// Assert
	var expectedSortedList *ListNode
	assert.Equal(t, expectedSortedList, sortedList)
}

func TestWhenList1NilThenReturnList2(t *testing.T) {
	// Arrange
	var list1 *ListNode
	list2 := &ListNode{
		Val: 1,
	}

	// Act
	sortedList := mergeTwoLists(list1, list2)

	// Assert
	assert.Equal(t, list2, sortedList)
}

func TestWhenList2NilThenReturnList1(t *testing.T) {
	// Arrange
	var list2 *ListNode
	list1 := &ListNode{
		Val: 1,
	}

	// Act
	sortedList := mergeTwoLists(list1, list2)

	// Assert
	assert.Equal(t, list1, sortedList)
}

func TestWhenAllList1GreaterThanList2(t *testing.T) {
	// Arrange
	list1 := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 6,
			},
		},
	}
	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	// Act
	sortedList := mergeTwoLists(list1, list2)

	// Assert
	expectedSortedList := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
						},
					},
				},
			},
		},
	}
	assert.Equal(t, expectedSortedList, sortedList)
}

func TestWhenAllList1LessThanList2(t *testing.T) {
	// Arrange
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	list2 := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 6,
			},
		},
	}

	// Act
	sortedList := mergeTwoLists(list1, list2)

	// Assert
	expectedSortedList := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
						},
					},
				},
			},
		},
	}
	assert.Equal(t, expectedSortedList, sortedList)
}

func TestWhenList1AndList2AreInterleaved(t *testing.T) {
	// Arrange
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 5,
			},
		},
	}
	list2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 6,
			},
		},
	}

	// Act
	sortedList := mergeTwoLists(list1, list2)

	// Assert
	expectedSortedList := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
						},
					},
				},
			},
		},
	}
	assert.Equal(t, expectedSortedList, sortedList)
}
