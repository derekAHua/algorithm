package T000_

// @Author: Derek
// @Description:
// @Date: 2023/9/9 20:26
// @Version 1.0

func minPathSum(grid [][]int) (ret int) {

	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}
	dp[0][0] = grid[0][0]

	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if j == 0 && i == 0 {
				continue
			}
			if j == 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
				continue
			}
			if i == 0 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
				continue
			}
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}

	return dp[len(grid)-1][len(grid[0])-1]
}
