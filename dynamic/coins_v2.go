package dynamic

// 1,2,5
func changeV1(amount int, coins []int) int {
	dp := make([][]int, len(coins)+1)
	for k := range dp {
		dp[k] = make([]int, amount+1)
		dp[k][0] = 1
	}

	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			}

		}
	}

	return dp[len(coins)][amount]
}
