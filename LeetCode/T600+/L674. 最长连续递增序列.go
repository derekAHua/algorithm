package T600_

// https://leetcode.cn/problems/longest-continuous-increasing-subsequence/

func findLengthOfLCIS(nums []int) (ret int) {
	ret = 1
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
		if ret < dp[i] {
			ret = dp[i]
		}
	}

	return
}

func findLengthOfLCISR3(nums []int) (ret int) {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	ret = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}

		if dp[i] > ret {
			ret = dp[i]
		}
	}

	return
}

func findLengthOfLCISR2(nums []int) (ret int) {
	l, r := 0, 1
	for ; r < len(nums); r++ {
		if nums[r] > nums[r-1] {
			continue
		}

		if ret < r-l {
			ret = r - l
		}
		l = r
	}

	if ret < r-l {
		ret = r - l
	}
	return
}

func findLengthOfLCISR1(nums []int) (ret int) {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	ret = 1

	for i := 1; i < len(dp); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}

		if dp[i] > ret {
			ret = dp[i]
		}
	}

	return
}
