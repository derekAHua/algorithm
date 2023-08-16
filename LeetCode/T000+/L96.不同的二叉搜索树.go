package T000_

// https://leetcode.cn/problems/unique-binary-search-trees/

func numTrees(n int) (ret int) {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i < len(dp); i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]
}

func numTreesR1(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i < len(dp); i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]
}
