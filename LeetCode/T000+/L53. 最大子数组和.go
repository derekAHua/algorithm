package T000_

// https://leetcode.cn/problems/maximum-subarray/

func maxSubArray(nums []int) (ret int) {
	ret = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > ret {
			ret = nums[i]
		}
	}
	return
}

func maxSubArrayR2(nums []int) (ret int) {
	ret = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > ret {
			ret = nums[i]
		}
	}
	return
}

func maxSubArrayR1(nums []int) (ret int) {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	ret = nums[0]

	for i := 1; i < len(dp); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		ret = max(ret, dp[i])
	}

	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
