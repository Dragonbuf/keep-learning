package trace

//def backtrack(路径, 选择列表):
//    if 满足结束条件:
//        result.add(路径)
//        return
//
//    for 选择 in 选择列表:
//        做选择
//        backtrack(路径, 选择列表)
//        撤销选择

//给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选择一个符号添加在前面。
// todo
//返回可以使最终数组和为目标数 S 的所有添加符号的方法数。
//输入：nums: [1, 1, 1, 1, 1], S: 3 输出：5

var resV1 [][]int

// 全排列 [1,2,4]
func permuteV1(nums []int) [][]int {
	var trace []int

	backTrackV1(nums, trace)
	return resV1
}

func backTrackV1(nums []int, trace []int) {
	if len(nums) == len(trace) {
		resV1 = append(resV1, trace)
		return
	}

	for _, v := range nums {
		if containsV1(trace, v) {
			continue
		}
		trace = append(trace, v)
		backTrackV1(nums, trace)
		trace = trace[:len(trace)-1]
	}
}

func containsV1(nums []int, target int) bool {
	for _, v := range nums {
		if v == target {
			return true
		}
	}

	return false
}

var resV2 [][]int

// 全排列 [1,2,4]
func permuteV2(nums []int) [][]int {
	var trace []int

	backTrackV1(nums, trace)
	return resV1
}

func backTrackV2(nums []int, trace []int) {
	if len(nums) == len(trace) {
		resV1 = append(resV1, trace)
		return
	}

	for _, v := range nums {
		if containsV1(trace, v) {
			continue
		}
		trace = append(trace, v)
		backTrackV1(nums, trace)
		trace = trace[:len(trace)-1]
	}
}
