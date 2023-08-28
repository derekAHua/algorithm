package T500_

// https://leetcode.cn/problems/longest-palindromic-subsequence/

func longestPalindromeSubseq(s string) (ret int) {
	size := len(s)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int, size)
		dp[i][i] = 1
	}
	ret = 1
	for i := len(s) - 1; i >= 0; i-- {
		for r := i + 1; r < len(s); r++ {
			if s[i] == s[r] {
				dp[i][r] = dp[i+1][r-1] + 2
			} else {
				dp[i][r] = max(dp[i][r-1], dp[i+1][r])
			}
		}
	}

	return dp[0][size-1]
}

func longestPalindromeSubseqR1(s string) int {
	size := len(s)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int, size)
		dp[i][i] = 1
	}

	for i := size - 1; i >= 0; i-- {
		for j := i + 1; j < size; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
		}
	}

	return dp[0][size-1]
}
