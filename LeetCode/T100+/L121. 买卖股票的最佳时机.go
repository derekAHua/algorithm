package T100_

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/

func maxProfit_(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dp := make([][2]int, len(prices))
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	return dp[len(dp)-1][1]
}

func maxProfit_R1(prices []int) (ret int) {
	if len(prices) <= 1 {
		return 0
	}

	min := prices[0]
	for i := 1; i < len(prices); i++ {
		min = mi(min, prices[i])
		ret = ma(ret, prices[i]-min)
	}
	return
}

func maxProfit_R2(prices []int) int {
	length := len(prices)
	if length == 0 {
		return 0
	}

	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = -prices[0] // 持有股票
	dp[0][1] = 0          // 不持有股票
	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[length-1][1]
}

func mi(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func ma(x, y int) int {
	if x < y {
		return y
	}
	return x
}
