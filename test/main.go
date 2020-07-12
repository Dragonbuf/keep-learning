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
