package T1000_

// https://leetcode.cn/problems/last-stone-weight-ii/

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, v := range stones {
		sum += v
	}
	target := sum >> 1
	dp := make([]int, target+1)
	for _, v := range stones {
		for i := len(dp) - 1; i >= v; i-- {
			if dp[i] < v+dp[i-v] {
				dp[i] = v + dp[i-v]
			}
		}
	}
	return sum - dp[len(dp)-1]<<1
}

func lastStoneWeightIIR3(stones []int) int {
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
