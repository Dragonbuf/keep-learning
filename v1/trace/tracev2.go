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

var resFindV1 int

func findTargetSumWaysV1(nums []int, target int) int {
	var trace int
	backTrackFindV1(nums, trace, 0, target)
	return resFindV1
}

func backTrackFindV1(nums []int, trace int, start int, target int) {
	if start == len(nums) {
		if trace == target {
			resFindV1++
		}
		return
	}

	chooses := []int{+1, -1}
	for _, choose := range chooses {
		trace += nums[start] * choose
		backTrackFindV1(nums, trace, start+1, target)
		trace -= nums[start] * choose
	}
}

// 会有重复子问题,比如 start:trace  会被进行重复计算
func backTrackFindV2(nums []int, trace int, start int, target int) {
	if start == len(nums) {
		if trace == target {
			resFindV1++
		}
		return
	}

	chooses := []int{+1, -1}
	for _, choose := range chooses {
		trace += nums[start] * choose
		backTrackFindV1(nums, trace, start+1, target)
		trace -= nums[start] * choose
	}
}

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

//4 toto N 皇后问题
var resV2 [][]string

func solveNQueens(n int) [][]string {
	return nil
}

func backTrackV2() {

}
