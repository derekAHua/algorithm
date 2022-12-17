package T000_

// https://leetcode.cn/problems/unique-paths/

func uniquePaths(m int, n int) int {
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
