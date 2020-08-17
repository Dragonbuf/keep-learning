package dynamic

//10 9 2 5 3 7 101 18
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for k := range dp {
		dp[k] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	var res int
	for j := 0; j < len(dp); j++ {
		res = max(res, dp[j])
	}

	return res
}

// 使用二分 + 纸牌思想解决问题
func lengthOfLISV2(nums []int) int {

	// 栈顶的数值
	tops := make([]int, len(nums))
	// 栈的数量
	n := 0

	for i := 0; i < len(nums); i++ {

		left := 0
		right := n
		mid := left + (right-left)/2

		// 搜索左边界的二分查找
		for left < right {
			if tops[mid] > nums[i] {
				right = mid
			} else if tops[mid] < nums[i] {
				left = mid + 1
			} else {
				right = mid
			}
		}

		// 没找到合适的
		if left == n {
			n++
		}
		tops[left] = nums[i]
	}

	return n
}
