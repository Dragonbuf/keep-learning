package dynamic

func longestCommonSubsequence(str1, str2 string) int {
	dp := make([][]int, len(str1)+1)
	for k := range dp {
		dp[k] = make([]int, len(str2)+1)
	}

	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(str1)][len(str2)]
}

func longestPalindromeSubseq(s string) int {
	size := len(s)
	dp := make([][]int, size)
	for k := range dp {
		dp[k] = make([]int, size)
	}

	for i := size - 1; i >= 0; i-- {
		for j := i + 1; j <= size-1; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
		}
	}

	return dp[0][size-1]
}
