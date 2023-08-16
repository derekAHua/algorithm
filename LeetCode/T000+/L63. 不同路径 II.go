package T000_

// https://leetcode.cn/problems/unique-paths-ii/

func uniquePathsWithObstacles(obstacleGrid [][]int) (ret int) {
	if len(obstacleGrid) == 0 || obstacleGrid[0][0] == 1 {
		return
	}

	dp := make([][]int, len(obstacleGrid))
	for i := range dp {
		dp[i] = make([]int, len(obstacleGrid[i]))
	}

	dp[0][0] = 1
	for i := 1; i < len(dp[0]); i++ {
		if obstacleGrid[0][i] == 1 {
			dp[0][i] = 0
			continue
		}
		dp[0][i] = dp[0][i-1]
	}
	for i := 1; i < len(dp); i++ {
		if obstacleGrid[i][0] == 1 {
			dp[i][0] = 0
			continue
		}
		dp[i][0] = dp[i-1][0]
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}

func uniquePathsWithObstaclesR1(obstacleGrid [][]int) int {
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
