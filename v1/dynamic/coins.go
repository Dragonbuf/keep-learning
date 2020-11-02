package dynamic

import (
	"math"
)

//：给你 k 种面值的硬币，面值分别为 c1, c2 ... ck，每种硬币的数量无限，再给一个总金额 amount，
//问你最少需要几枚硬币凑出这个金额，如果不可能凑出，算法返回 -1 。算法的函数签名如下：
//比如说 k = 3，面值分别为 1，2，5，总金额 amount = 11。那么最少需要 3 枚硬币凑出，即 11 = 5 + 5 + 1。
func coinChangev1(coins []int, amount int) int {
	return dpv1(amount, coins)
}

func dpv1(n int, coins []int) int {
	if n < 0 {
		return -1
	}

	if n == 0 {
		return 0
	}

	res := math.MaxInt64
	for _, coin := range coins {
		subproblem := dpv1(n-coin, coins)
		if subproblem == -1 {
			continue
		}
		res = min(res, subproblem+1)
	}

	if res == math.MaxInt64 {
		return -1
	} else {
		return res
	}
}

//带备忘录的递归
func coinChangev2(coins []int, amount int) int {
	m := make(map[int]int, amount+1)
	return dpv2(amount, coins, m)
}

func dpv2(n int, coins []int, m map[int]int) int {
	if n < 0 {
		return -1
	}

	if n == 0 {
		return 0
	}

	if v, ok := m[n]; ok {
		return v
	}

	res := math.MaxInt64
	for _, coin := range coins {
		subproblem := dpv1(n-coin, coins)
		if subproblem == -1 {
			continue
		}
		res = min(res, subproblem+1)
	}

	if res == math.MaxInt64 {
		return -1
	} else {
		m[n] = res
		return res
	}
}

//dp 数组的迭代解法

func dpv3(amount int, coins []int) int {

	dp := make([]int, amount+1)
	for k := range dp {
		dp[k] = amount + 1
	}

	dp[0] = 0

	for i := 0; i < len(dp); i++ {
		for _, coin := range coins {
			if coin > i {
				continue
			}

			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}
