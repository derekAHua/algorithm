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
