package dynamic

func canPartition(nums []int) bool {
	var sum int
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	dp := make([][]bool, len(nums)+1)
	for k := range dp {
		dp[k] = make([]bool, target+1)
		dp[k][0] = true
	}

	for i := 1; i <= len(nums); i++ {
		for j := 1; i <= target; j++ {
			if nums[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}

	return dp[len(nums)][target]
}
