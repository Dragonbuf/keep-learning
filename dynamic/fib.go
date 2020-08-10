package dynamic

func fibv1(n int) int {
	if n <= 2 {
		return 1
	}

	return fibv1(n-1) + fibv1(n-2)
}

// 带备忘录的 hash
func fibv2(n int) int {
	if n <= 0 {
		return 0
	}

	m := make(map[int]int, n+1)
	return v2Helper(m, n)
}

func v2Helper(m map[int]int, n int) int {
	if n <= 2 {
		return 1
	}

	if v, ok := m[n]; ok {
		return v
	}
	m[n] = v2Helper(m, n-1) + v2Helper(m, n-2)
	return m[n]
}

//自底向上

func fibv3(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// 压缩状态
func fibv4(n int) int {

	if n <= 2 {
		return 1
	}

	dp_1 := 1
	dp_2 := 1
	dp_n := 0
	for i := 3; i <= n; i++ {
		dp_n = dp_2 + dp_1
		dp_1 = dp_2
		dp_2 = dp_n
	}

	return dp_n
}
