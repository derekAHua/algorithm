package T400_

import "math"

// https://leetcode.cn/problems/target-sum/

func findTargetSumWays(nums []int, target int) (ret int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if int(math.Abs(float64(target))) > sum {
		return 0
	}
	if (sum+target)%2 == 1 {
		return
	}

	x := (target + sum) / 2
	dp := make([]int, x+1)
	dp[0] = 1

	for _, v := range nums {
		for j := x; j >= v; j-- {
			dp[j] += dp[j-v]
		}
	}

	return dp[x]
}

func findTargetSumWaysR4(nums []int, target int) (ret int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if int(math.Abs(float64(target))) > sum {
		return 0
	}
	if (sum+target)%2 == 1 {
		return
	}

	x := (target + sum) / 2
	dp := make([]int, x+1)
	dp[0] = 1

	for _, v := range nums {
		for j := x; j >= v; j-- {
			dp[j] += dp[j-v]
		}
	}

	return dp[x]
}

func findTargetSumWaysR3(nums []int, target int) (ret int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if int(math.Abs(float64(target))) > sum {
		return 0
	}
	if (sum+target)%2 == 1 {
		return 0
	}
	x := (target + sum) / 2
	dp := make([]int, x+1)
	dp[0] = 1
	for _, v := range nums {
		for j := x; j >= v; j-- {
			dp[j] += dp[j-v]
		}
	}

	return dp[x]
}

func findTargetSumWaysR2(nums []int, target int) (ret int) {
	var f func(idx, sum int)

	f = func(idx, sum int) {
		if idx == len(nums) {
			if sum == target {
				ret++
			}
			return
		}

		f(idx+1, sum+nums[idx])
		f(idx+1, sum-nums[idx])
	}

	f(0, 0)
	return
}

func findTargetSumWaysR1(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if int(math.Abs(float64(target))) > sum {
		return 0
	}
	if (sum+target)%2 == 1 {
		return 0
	}

	bag := (sum + target) / 2
	dp := make([]int, bag+1)
	dp[0] = 1
	for _, v := range nums {
		for i := bag; i >= v; i-- {
			dp[i] += dp[i-v]
		}
	}

	return dp[bag]
}
