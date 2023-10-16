package T600_

// https://leetcode.cn/problems/max-area-of-island/

func maxAreaOfIsland(grid [][]int) (ret int) {
	var dfs func(int, int) int
	dfs = func(i, j int) (cnt int) {
		if i < 0 || j < 0 || i == len(grid) || j == len(grid[0]) || grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0
		cnt++
		cnt += dfs(i, j+1)
		cnt += dfs(i, j-1)
		cnt += dfs(i+1, j)
		cnt += dfs(i-1, j)
		return
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				if r := dfs(i, j); r > ret {
					ret = r
				}
			}
		}
	}

	return
}
