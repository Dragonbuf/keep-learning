package dynamic

import "sort"

//给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。
//输入: [ [1,2], [1,2], [1,2] ]
//
//输出: 2
//
//解释: 你需要移除两个 [1,2] 来使剩下的区间没有重叠。
func intervalSchedule(intvs [][]int) int {

	// 按照 end 升序排列
	sort.Slice(intvs, func(i, j int) bool {
		return intvs[i][1] < intvs[j][1]
	})

	x := intvs[0]
	res := 0

	for i := 1; i < len(intvs); i++ {
		if intvs[i][0] >= x[1] {
			x = intvs[i]
			res++
		}
	}

	return len(intvs) - res - 1
}
