package T200_

// @Author: Derek
// @Description:
// @Date: 2023/9/3 15:10
// @Version 1.0

func numIslands(grid [][]byte) (ret int) {
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i == len(grid) || j == len(grid[i]) || grid[i][j] == '0' {
			return
		}

		grid[i][j] = '0'

		dfs(i, j+1)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i-1, j)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				ret++
				dfs(i, j)
			}
		}
	}
	return
}
