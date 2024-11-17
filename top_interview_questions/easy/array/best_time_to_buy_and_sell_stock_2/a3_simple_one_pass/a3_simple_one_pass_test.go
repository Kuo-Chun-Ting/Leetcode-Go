package a3_simple_one_pass

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenNums1EqualsToNums2AndLengthIsOne(t *testing.T) {
	// Arrange
	stocks := []int{1}

	// Act
	result := maxProfit(stocks)

	// Assert
	expectedMaxProfit := 0
	assert.Equal(t, expectedMaxProfit, result)
}
