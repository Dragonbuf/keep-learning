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
