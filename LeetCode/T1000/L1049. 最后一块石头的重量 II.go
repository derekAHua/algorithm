package T1000

// https://leetcode.cn/problems/last-stone-weight-ii/

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, v := range stones {
		sum += v
	}

	target := sum / 2
	dp := make([]int, target+1)
	for _, v := range stones {
		for j := target; j >= v; j-- {
			if dp[j] > dp[j-v]+v {
				dp[j] = dp[j-v] + v
			}
		}
	}

	return sum - dp[target]*2
}

func lastStoneWeightIIR2(stones []int) int {
	sum := 0
	for _, v := range stones {
		sum += v
	}

	t := sum / 2

	dp := make([]int, t+1)
	for _, v := range stones {
		for j := t; j >= v; j-- {
			if dp[j] < dp[j-v]+v {
				dp[j] = dp[j-v] + v
			}
		}
	}

	return sum - dp[t]*2
}

func lastStoneWeightIIR1(stones []int) int {

	sum := 0
	for _, v := range stones {
		sum += v
	}

	target := sum / 2
	dp := make([]int, target+1)
	for _, v := range stones {
		for i := len(dp) - 1; i >= v; i-- {
			dp[i] = max(dp[i], dp[i-v]+v)
		}
	}

	return sum - 2*dp[target]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
