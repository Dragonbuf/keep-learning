package dynamic

//给你 k 种面值的硬币，面值分别为 c1, c2 ... ck，
//每种硬币的数量无限，再给一个总金额 amount，
//问你最少需要几枚硬币凑出这个金额，如果不可能凑出，算法返回 -1 。算法的函数签名如下：

// coins 中是可选硬币面值，amount 是目标金额
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}

	for i := 0; i < len(dp); i++ {
		// todo
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

//在 n 个物品中挑选若干物品装入背包，最多能装多满？假设背包的大小为 m，每个物品的大小为 A[i]
//样例 1:
//	输入:  [3,4,8,5], backpack size=10
//	输出:  9
//
//样例 2:
//	输入:  [2,3,5,7], backpack size=12
//	输出:  12
func backPack(m int, A []int) int {
	dp := make([]int, m+1)
	for k := range dp {
		dp[k] = m + 1
	}
	dp[0] = 0
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(A); j++ {
			if A[j] > i {
				continue
			}
			dp[i] = min(dp[i], dp[i-A[j]]+1)
		}
	}

	if dp[m] == m+1 {
		return -1
	}

	return dp[m]
}

func min(a, b int) int {
	if a > b {
		return a
	}
	return b
}
