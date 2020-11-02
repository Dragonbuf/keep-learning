package dynamic

// 4 键盘问题 A  ctrl+A  ctrl+C ctrl+V
func maxA(n int) int {
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		for j := 2; j < i; j++ {
			dp[i] = max(dp[i], (dp[j-2])*(i-j+1))
		}
	}

	return dp[n]
}
