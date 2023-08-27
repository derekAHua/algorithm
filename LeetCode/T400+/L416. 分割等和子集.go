package T400_

import (
	"sort"
)

// https://leetcode.cn/problems/partition-equal-subset-sum/

func canPartition(nums []int) (ret bool) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 1 {
		return
	}

	target := sum / 2
	dp := make([]int, target+1)
	for _, v := range nums {
		for i := target; i >= v; i-- {
			if dp[i] < dp[i-v]+v {
				dp[i] = dp[i-v] + v
			}
		}
	}

	return dp[target] == target
}

func canPartitionR3(nums []int) (ret bool) {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 == 1 {
		return
	}

	target := sum / 2
	dp := make([]int, target+1)

	for _, v := range nums {
		for j := target; j >= v; j-- {
			if dp[j] < dp[j-v]+v {
				dp[j] = dp[j-v] + v
			}
		}
	}

	return dp[target] == target
}

func canPartitionR2(nums []int) (ret bool) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 1 {
		return false
	}

	target := sum / 2
	dp := make([]int, target+1)

	for _, v := range nums {
		for j := target; j >= v; j-- {
			if dp[j] < dp[j-v]+v {
				dp[j] = dp[j-v] + v
			}
		}
	}

	return dp[target] == target
}

func canPartitionR1(nums []int) (ret bool) {
	sort.Ints(nums)

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	we := sum % 2
	if we == 1 {
		return false
	}

	tar := sum / 2
	dp := make([]int, tar+1)
	for _, v := range nums {
		for i := tar; i >= v; i-- {
			if dp[i] < dp[i-v]+v {
				dp[i] = dp[i-v] + v
			}
		}
	}

	return dp[tar] == tar
}
