package T1000_

// @Author: Derek
// @Description:
// @Date: 2023/8/26 19:52
// @Version 1.0

func validPath2(n int, edges [][]int, source int, destination int) bool {
	vis := make([]bool, n)
	g := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	var dfs func(int) bool
	dfs = func(i int) bool {
		if i == destination {
			return true
		}
		vis[i] = true

		for _, j := range g[i] {
			if !vis[j] && dfs(j) {
				return true
			}
		}
		return false
	}
	return dfs(source)
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	vis := make([]bool, n)
	g := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	var f func(cur int) bool
	f = func(cur int) bool {
		if cur == destination {
			return true
		}

		vis[cur] = true
		for _, v := range g[cur] {
			if !vis[v] && f(v) {
				return true
			}
		}

		return false
	}

	return f(source)
}
