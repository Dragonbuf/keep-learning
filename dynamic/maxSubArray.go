package dynamic

import "math"

//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//-2,1,-3,4,-1,2,1,-5,4
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
	}

	res := math.MinInt64
	for _, v := range dp {
		res = max(res, v)
	}
	return res
}
