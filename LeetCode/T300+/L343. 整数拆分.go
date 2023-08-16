package T300_

// https://leetcode.cn/problems/integer-break/

func integerBreak(n int) (ret int) {
	dp := make([]int, n+1)
	for i := 2; i < len(dp); i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}

	return dp[n]
}

func integerBreakR1(n int) int {
	dp := make([]int, n+1)
	for i := 2; i < len(dp); i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}

	return dp[n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
