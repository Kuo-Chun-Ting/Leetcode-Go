package a2_peak_valley_approach

// Check twice
func maxProfit(prices []int) int {
	i := 0
	maxProfit := 0

	for i < len(prices)-1 {
		for i < len(prices)-1 && prices[i] >= prices[i+1] {
			i++
		}
		valley := prices[i]

		for i < len(prices)-1 && prices[i] <= prices[i+1] {
			i++
		}
		peak := prices[i]

		maxProfit += peak - valley
	}

	return maxProfit
}
