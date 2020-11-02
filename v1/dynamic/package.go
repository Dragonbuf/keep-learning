package dynamic

//给你一个可装载重量为 W 的背包和 N 个物品，
//每个物品有重量和价值两个属性。其中第i个物品的重量为wt[i]，价值为val[i]
//，现在让你用这个背包装物品，最多能装的价值是多少？
//N = 3, W = 4
//wt = [2, 1, 3]
//val = [4, 2, 3]

//算法返回 6，选择前两件物品装进背包，总重量 3 小于W，可以获得最大价值 6。

// 状态， 选择， dp 数组定义

func knapsack(N, W int, wt, val []int) int {
	dp := make([][]int, N+1)
	for k := range dp {
		dp[k] = make([]int, W+1)
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= W; j++ {
			if j-wt[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-wt[i-1]]+val[i-1])
			}
		}
	}

	return dp[N][W]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
