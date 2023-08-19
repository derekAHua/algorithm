package orther

// @Author: Derek
// @Description:
// @Date: 2023/8/18 22:29
// @Version 1.0

// https://leetcode.cn/problems/coin-lcci/description/

func waysToChange(n int) (ret int) {
	dp := make([]int, n+1)
	dp[0] = 1
	coins := []int{1, 5, 10, 25}

	for i := 0; i < 4; i++ {
		for j := 1; j <= n; j++ {
			if j-coins[i] >= 0 {
				dp[j] += dp[j-coins[i]]
			}
		}
	}

	return dp[n] % 1000000007
}
