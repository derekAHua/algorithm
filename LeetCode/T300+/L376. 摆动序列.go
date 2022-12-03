package T300_

// https://leetcode.cn/problems/wiggle-subsequence/

func wiggleMaxLength(nums []int) (ret int) {
	if len(nums) < 2 {
		return len(nums)
	}

	f := true
	ret++
	start := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			f = nums[i]-nums[i-1] < 0
			start = i + 1
			break
		}
	}
	if start == 0 {
		return
	}
	ret++

	for i := start; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > nums[i-1] && f {
			f = false
			ret++
		}
		if nums[i] < nums[i-1] && !f {
			f = true
			ret++
		}
	}

	return
}
