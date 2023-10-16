package T1000_

// https://leetcode.cn/problems/number-of-enclaves/description/

func numEnclaves(grid [][]int) (ret int) {
	var dfs func(i, j int) (cnt int, is bool)
	dfs = func(i, j int) (cnt int, is bool) {
		if i < 0 || j < 0 || i == len(grid) || j == len(grid[0]) {
			is = true
			return
		}

		if grid[i][j] == 0 {
			return
		}

		grid[i][j] = 0
		v1, is1 := dfs(i, j+1)
		v2, is2 := dfs(i, j-1)
		v3, is3 := dfs(i+1, j)
		v4, is4 := dfs(i-1, j)

		return 1 + v1 + v2 + v3 + v4, is || is1 || is2 || is3 || is4
	}

	for i, v := range grid {
		for j := range v {
			if grid[i][j] == 1 {
				r, is := dfs(i, j)
				if !is {
					ret += r
				}
			}
		}
	}
	return
}
