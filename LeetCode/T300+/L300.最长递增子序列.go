package T300_

// https://leetcode.cn/problems/longest-increasing-subsequence/

//给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//
//子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

func lengthOfLIS(nums []int) (ret int) {
	dp := make([]int, len(nums))
	dp[0] = 1
	ret = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		if ret < dp[i] {
			ret = dp[i]
		}
	}

	return
}

func lengthOfLISR2(nums []int) (ret int) {
	if len(nums) == 0 {
		return
	}
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	ret = 1

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] <= nums[j] {
				continue
			}
			if dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}

		if dp[i] > ret {
			ret = dp[i]
		}
	}

	return
}

func lengthOfLISR1(nums []int) (ret int) {
	dp := make([]int, len(nums))

	for i := range dp {
		dp[i] = 1
	}

	ret = 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > ret {
			ret = dp[i]
		}
	}
	return
}
