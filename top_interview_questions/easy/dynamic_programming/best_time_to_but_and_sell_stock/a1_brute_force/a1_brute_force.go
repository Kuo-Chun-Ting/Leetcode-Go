package a1_brute_force

func maxProfit(prices []int) int {
	maxProfit := 0

	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if profit := prices[j] - prices[i]; profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}
