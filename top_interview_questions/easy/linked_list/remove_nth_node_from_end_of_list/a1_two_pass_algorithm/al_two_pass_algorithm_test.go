package a1_two_pass_algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountLengthWhenNodeIsNilThenReturn0(t *testing.T) {
	// Arrange
	var node *ListNode = nil

	// Act
	length := countLength(node)

	// Assert
	assert.Equal(t, 0, length)
}

func TestCountLengthWhenOneNodeThenReturn1(t *testing.T) {
	// Arrange
	node := &ListNode{Val: 1}

	// Act
	length := countLength(node)

	// Assert
	assert.Equal(t, 1, length)
}

func TestCountLengthWhenTwoNodeThenReturn2(t *testing.T) {
	// Arrange
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}

	// Act
	length := countLength(node)

	// Assert
	assert.Equal(t, 2, length)
}

func TestCountLengthWhenThreeNodeThenReturn3(t *testing.T) {
	// Arrange
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	// Act
	length := countLength(node)

	// Assert
	assert.Equal(t, 3, length)
}

func TestRemoveNthFromEndWhenSizeOneReturnNil(t *testing.T) {
	// Arrange
	head := &ListNode{Val: 1}
	n := 1

	// Act
	result := removeNthFromEnd(head, n)

	// Assert
	assert.Nil(t, result)
}

func TestRemoveNthFromEndWhenRemoveFirstNodeThenResultShouldBeCorrect(t *testing.T) {
	// Arrange
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	n := 3

	// Act
	result := removeNthFromEnd(head, n)

	// Assert
	expectedResult := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
		},
	}
	assert.Equal(t, expectedResult, result)
}

func TestRemoveNthFromEndWhenRemoveLastNodeThenResultShouldBeCorrect(t *testing.T) {
	// Arrange
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	n := 1

	// Act
	result := removeNthFromEnd(head, n)

	// Assert
	expectedResult := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	assert.Equal(t, expectedResult, result)
}

func TestRemoveNthFromEndWhenRemoveMiddleNodeThenResultShouldBeCorrect(t *testing.T) {
	// Arrange
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	n := 2

	// Act
	result := removeNthFromEnd(head, n)

	// Assert
	expectedResult := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
		},
	}
	assert.Equal(t, expectedResult, result)
}
