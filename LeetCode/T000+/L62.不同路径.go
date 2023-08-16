package T000_

// https://leetcode.cn/problems/unique-paths/

func uniquePaths(m int, n int) (ret int) {
	if m == 0 || n == 0 {
		return
	}

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := range dp[0] {
		dp[0][i] = 1
	}
	for i := 1; i < len(dp); i++ {
		dp[i][0] = 1
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

func uniquePathsR1(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 1; i < n; i++ {
		dp[0][i] = 1
	}
	for j := 0; j < m; j++ {
		dp[j][0] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[j][i] = dp[j-1][i] + dp[j][i-1]
		}
	}

	return dp[m-1][n-1]
}
