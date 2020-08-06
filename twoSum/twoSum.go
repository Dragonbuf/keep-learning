package twoSum

import (
	"sort"
)

//如果假设输入一个数组 nums 和一个目标和 target，
//请你返回 nums 中能够凑出 target 的两个元素的值，
//比如输入 nums = [1,3,5,6], target = 9，
//那么算法返回两个元素 [3,6]。
//可以假设只有且仅有一对儿元素可以凑出 target。

func twoSum(nums []int, target int) []int {
	sort.Ints(nums)
	left := 0
	right := len(nums) - 1
	for left < right {
		sum := nums[left] + nums[right]
		if sum > target {
			right--
		} else if sum < target {
			left++
		} else if sum == target {
			return []int{nums[left], nums[right]}
		}
	}

	return []int{left, right}
}

//nums 中可能有多对儿元素之和都等于 target，请你的算法返回所有和为 target 的元素对儿，其中不能出现重复。
//比如说输入为 nums = [1,3,1,2,2,3], target = 4，那么算法返回的结果就是：[[1,3],[2,2]]。
func twoSum2(nums []int, target int) (res [][]int) {
	sort.Ints(nums)
	left := 0
	right := len(nums) - 1
	for left < right {
		sum := nums[left] + nums[right]
		l := left
		r := right
		if sum > target {
			for left < right && nums[r] == nums[right] {
				right--
			}
		} else if sum < target {
			for left < right && nums[l] == nums[left] {
				left++
			}
		} else if sum == target {
			res = append(res, []int{nums[left], nums[right]})
			for left < right && nums[l] == nums[left] {
				left++
			}
			for left < right && nums[r] == nums[right] {
				right--
			}
		}
	}

	return
}
