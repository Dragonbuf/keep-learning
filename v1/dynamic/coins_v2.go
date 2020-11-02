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

func changeV2(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] >= 0 {
				dp[j] += dp[j-coins[i-1]]
			}
		}
	}

	return dp[amount]
}

// 戳气球问题
//https://mp.weixin.qq.com/s/I0yo0XZamm-jMpG-_B3G8g
func maxCoins(nums []int) int {

	n := len(nums)
	points := make([]int, n+2)
	points[0] = 1
	points[n+1] = 1
	for k := range nums {
		points[k+1] = nums[k]
	}

	dp := make([][]int, n+2)
	for k := range dp {
		dp[k] = make([]int, n+2)
	}

	for i := n; i >= 0; i-- {
		for j := i + 1; j <= n+1; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+points[i]*points[j]*points[k])
			}
		}
	}

	return dp[0][n+1]
}
