package LeetCode

func maximizeTheProfit(n int, offers [][]int) int {
	type pair struct{ start, gold int }

	groups := make([][]pair, n)
	for _, offer := range offers {
		start, end, gold := offer[0], offer[1], offer[2]
		groups[end] = append(groups[end], pair{start, gold})
	}

	dp := make([]int, n+1)

	for end, ll := range groups {
		dp[end+1] = dp[end]
		for _, v := range ll {
			cur := v.gold + dp[v.start]
			if cur > dp[end+1] {
				dp[end+1] = cur
			}
		}
	}

	return dp[n]
}
