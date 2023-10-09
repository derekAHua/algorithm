package T1000_

// @Author: Derek
// @Description:
// @Date: 2023/9/29 09:55
// @Version 1.0

func tribonacci(n int) (ret int) {
	if n == 0 {
		return
	}
	if n < 3 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}

	return dp[n]
}
