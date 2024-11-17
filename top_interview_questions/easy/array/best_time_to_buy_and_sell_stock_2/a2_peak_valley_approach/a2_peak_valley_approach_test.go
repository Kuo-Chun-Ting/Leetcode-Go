package a2_peak_valley_approach

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenOnlyOneDayThenReturn0(t *testing.T) {
	// Arrange
	stocks := []int{7}

	// Act
	result := maxProfit(stocks)

	// Assert
	expectedMaxProfit := 0
	assert.Equal(t, expectedMaxProfit, result)
}

func TestWhenNoWayToMakeProfitThenReturn0(t *testing.T) {
	// Arrange
	stocks := []int{7, 6, 5, 4}

	// Act
	result := maxProfit(stocks)

	// Assert
	expectedMaxProfit := 0
	assert.Equal(t, expectedMaxProfit, result)
}

func TestWhenMaxProfitWayExistsThenReturnMaxProfit(t *testing.T) {
	// Arrange
	stocks := []int{7, 1, 5, 3, 6, 4}

	// Act
	result := maxProfit(stocks)

	// Assert
	expectedMaxProfit := 7
	assert.Equal(t, expectedMaxProfit, result)
}

func TestWhenConsecutiveHighOrLowThenReturnMaxProfit(t *testing.T) {
	// Arrange
	stocks := []int{7, 1, 3, 5, 4, 3, 6, 4}

	// Act
	result := maxProfit(stocks)

	// Assert
	expectedMaxProfit := 7
	assert.Equal(t, expectedMaxProfit, result)
}
