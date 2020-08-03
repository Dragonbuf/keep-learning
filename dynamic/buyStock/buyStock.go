package buyStock

import "math"

func buy1(prices []int) int {
	n := len(prices)
	dp_i_0 := 0
	dp_i_1 := math.MinInt64
	for i := 0; i < n; i++ {
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, -prices[i])
	}

	return dp_i_0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//todo https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
//todo https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/
//todo https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/
//todo https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
//todo https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
