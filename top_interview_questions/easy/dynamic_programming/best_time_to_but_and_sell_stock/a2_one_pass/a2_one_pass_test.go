package a2_one_pass

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenHasMaxProfit(t *testing.T) {
	// Arrange
	profits := []int{7, 1, 5, 3, 6, 4}
	// Act
	maxProfit := maxProfit(profits)

	// Assert
	expectedMaxProfit := 5
	assert.Equal(t, expectedMaxProfit, maxProfit)
}

func TestWhenNoMaxProfit(t *testing.T) {
	// Arrange
	profits := []int{7, 6, 4, 3, 1}
	// Act
	maxProfit := maxProfit(profits)

	// Assert
	expectedMaxProfit := 0
	assert.Equal(t, expectedMaxProfit, maxProfit)
}
