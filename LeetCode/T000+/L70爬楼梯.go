package T000_

// https://leetcode.cn/problems/climbing-stairs/

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2

	for i := 3; i < len(dp); i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func climbStairs2(n int) int {
	dp := make([]int, n+1)
	nums := []int{1, 2}
	dp[0] = 1
	for i := 1; i < len(dp); i++ {
		for _, v := range nums {
			if i >= v {
				dp[i] += dp[i-v]
			}
		}
	}

	return dp[n]
}
