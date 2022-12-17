package T000_

// https://leetcode.cn/problems/unique-paths-ii/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}
	for j := 0; j < m; j++ {
		if obstacleGrid[j][0] == 1 {
			break
		}
		dp[j][0] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if obstacleGrid[j][i] == 1 {
				continue
			}
			dp[j][i] += dp[j][i-1] + dp[j-1][i]
		}
	}

	return dp[m-1][n-1]
}
