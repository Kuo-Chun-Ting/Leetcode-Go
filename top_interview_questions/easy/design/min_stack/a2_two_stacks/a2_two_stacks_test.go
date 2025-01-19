package a2_two_stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenPushOnceThenCorrectMinAndTop(t *testing.T) {
	// Arrange
	stack := Constructor()

	// Act & Assert
	stack.Push(1)
	assert.Equal(t, 1, stack.GetMin())
	assert.Equal(t, 1, stack.Top())
}

func TestWhenMinValueChangesThenGetMinReflectsIt(t *testing.T) {
	// Arrange
	stack := Constructor()

	// Act & Assert
	stack.Push(2)
	assert.Equal(t, 2, stack.GetMin())
	assert.Equal(t, 2, stack.Top())

	stack.Push(1)
	assert.Equal(t, 1, stack.GetMin())
	assert.Equal(t, 1, stack.Top())

	stack.Pop()
	assert.Equal(t, 2, stack.GetMin())
	assert.Equal(t, 2, stack.Top())
}

func TestWhenMultiplePushAndPopThenCorrectMinAndTop(t *testing.T) {
	// Arrange
	stack := Constructor()

	// Act & Assert
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	assert.Equal(t, -3, stack.GetMin())

	stack.Pop()
	assert.Equal(t, 0, stack.Top())
	assert.Equal(t, -2, stack.GetMin())
}
