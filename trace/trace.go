package trace

//给定一个 没有重复 数字的序列，返回其所有可能的全排列。

var res [][]int

func permute(nums []int) [][]int {
	var track []int
	backtrack(nums, track)
	return res
}

func backtrack(nums, track []int) {
	// 终止条件
	if len(nums) == len(track) {
		res = append(res, track)
		return
	}

	for i := 0; i < len(nums); i++ {
		if contains(track, nums[i]) {
			continue
		}
		// 做选择
		track = append(track, nums[i])
		// 进入决策树
		backtrack(nums, track)
		// 撤销选择
		track = track[:len(track)-1]
	}

}

func contains(track []int, needle int) bool {
	for _, v := range track {
		if v == needle {
			return true
		}
	}
	return false
}

func subsetsWithDup(nums []int) [][]int {
	var track []int
	backtrack2(nums, 0, track)
	return res
}

func backtrack2(nums []int, start int, track []int) {
	res = append(res, track)
	for i := start; i < len(nums); i++ {
		track = append(track, nums[i])
		backtrack2(nums, i+1, track)
		track = track[:len(track)-1]
	}
}
