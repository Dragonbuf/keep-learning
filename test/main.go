package main

import (
	"fmt"
	"github.com/imroc/biu"
)

func main() {
	var a uint8
	a = 3
	// uint32 把最后一位取出来
	p := uint8(7)

	fmt.Println((a&1)<<p, biu.ToBinaryString((a & 1)), biu.ToBinaryString((a&1)<<p))

}

func minimumTotal(triangle [][]int) int {

	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}

	return triangle[0][0]
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func longestConsecutive(nums []int) int {

	uniqueMap := map[int]bool{}
	// 先初始化到 map 中
	for _, v := range nums {
		uniqueMap[v] = true
	}

	var max int

	for k, _ := range uniqueMap {
		if !uniqueMap[k-1] {
			num := 1
			current := k
			for uniqueMap[current+1] {
				num++
				current++
			}
			if num > max {
				max = num
			}
		}
	}

	return max
}

//给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//说明：每次只能向下或者向右移动一步。

//状态 State 存储小规模问题的结果
//方程 Function 怎么通过小的状态，来算大的状态
//初始化 Intialization 最极限的小状态是什么, 起点
//答案 Answer

func uniquePaths(m int, n int) int {
	//  f(m,n) = f(i-1,j)+f(i,j-1)
	f := make([][]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if f[i] == nil {
				f[i] = make([]int, n)
			}
			f[i][j] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			f[i][j] = f[i-1][j] + f[i][j-1]
		}
	}
	return f[m-1][n-1]
}

//有障碍物
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[m-1][n-1] == 1 || obstacleGrid[0][0] == 1 {
		return 0
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
			} else {

				// 额外考虑 左边没有数据，上方为0的，自己也必须为0
				if j == 0 && i-1 >= 0 && obstacleGrid[i-1][j] == 0 {
					obstacleGrid[i][j] = 0
					continue
				}

				// 考虑上面没有数据，左边为0的，自己也必须为0
				if i == 0 && j-1 >= 0 && obstacleGrid[i][j-1] == 0 {
					obstacleGrid[i][j] = 0
					continue
				}

				obstacleGrid[i][j] = 1
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				continue
			} else {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			}

		}
	}
	return obstacleGrid[m-1][n-1]
}

// 爬楼梯
func climbStairs(n int) int {
	if n <= 1 {
		return n
	}
	f := make([]int, n+1)

	f[1] = 1
	f[2] = 2
	for i := 3; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

//给定一个非负整数数组，你最初位于数组的第一个位置。
//数组中的每个元素代表你在该位置可以跳跃的最大长度。 判断你是否能够到达最后一个位置。
//[2,3,1,1,4]
func canJump(nums []int) bool {
	rightMost := 0
	for i := 0; i < len(nums); i++ {
		if i <= rightMost {
			rightMost = max(nums[i]+i, rightMost)
			if rightMost >= len(nums)-1 {
				return true
			}
		}
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func canJump2(nums []int) int {
	var end, maxPosition, step int
	//在遍历数组时，我们不访问最后一个元素，这是因为在访问最后一个元素之前，我们的边界一定大于等于最后一个位置，否则就无法跳到最后一个位置了。
	//如果访问最后一个元素，在边界正好为最后一个位置的情况下，我们会增加一次「不必要的跳跃次数」，因此我们不必访问最后一个元素。

	for i := 0; i < len(nums)-1; i++ {
		maxPosition = max(maxPosition, nums[i]+i)
		if end == i {
			end = maxPosition
			step++
		}
	}
	return step
}

//给定一个无序的整数数组，找到其中最长上升子序列的长度。
//[10,9,2,5,3,7,101,18]
//[2,3,7,101]，它的长度是 4。
func lengthOfLIS(nums []int) int {
	// dp[i] = max(dp[i])+1

	if len(nums) < 2 {
		return len(nums)
	}

	dp := make([]int, len(nums))
	dp[0] = 1

	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	var res int
	for _, v := range dp {
		res = max(res, v)
	}

	return res
}

//给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
//输入: s = "leetcode", wordDict = ["leet", "code"]
//输出: true
//解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。

//输入: s = "applepenapple", wordDict = ["apple", "pen"]
//输出: true
//解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
//     注意你可以重复使用字典中的单词。

//输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
//输出: false
