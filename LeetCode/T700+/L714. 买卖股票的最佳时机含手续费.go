package T700_

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/

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
