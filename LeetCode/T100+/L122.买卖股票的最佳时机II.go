package T100_

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/

func maxProfit(prices []int) (ret int) {
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ret += prices[i] - prices[i-1]
		}
	}
	return
}

func maxProfit2(prices []int) (ret int) {
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			ret += prices[i] - prices[i-1]
		}
	}
	return
}
