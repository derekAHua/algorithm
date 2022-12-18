package T700_

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/

func maxProfit2(prices []int, fee int) (ret int) {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][0] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
	}

	return dp[n-1][1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProfit(prices []int, fee int) (ret int) {
	var minBuy = prices[0]

	for i := 0; i < len(prices); i++ {
		if prices[i] < minBuy {
			minBuy = prices[i]
		}

		if prices[i]-minBuy <= fee {
			continue
		}

		ret += prices[i] - minBuy - fee
		minBuy = prices[i] - fee
	}

	return
}
