package _package

func canPartition(nums []int) bool {
	var sum int
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	sum /= 2

	dp := make([]bool, sum+1)
	dp[0] = true

	for i := 0; i < len(nums); i++ {
		for j := sum; j >= 0; j-- {
			if nums[i] <= j {
				dp[j] = dp[j] || dp[j-nums[i]]
			}
		}
	}

	return dp[sum]
}
