package T200_

// https://leetcode-cn.com/problems/minimum-size-subarray-sum/

/*
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。

示例：

输入：s = 7, nums = [2,3,1,2,4,3] 输出：2 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
*/

func minSubArrayLen(target int, nums []int) (ret int) {
	l := len(nums)
	if l == 0 {
		return 0
	}

	left, sum := 0, 0        // left：滑动窗口左指针
	ret = l + 1              // 初始化返回长度为l+1，目的是为了判断“不存在符合条件的子数组，返回0”的情况
	for i := 0; i < l; i++ { // i：滑动窗口右指针
		sum += nums[i]
		for sum >= target {
			subLength := i - left + 1
			if subLength < ret {
				ret = subLength
			}
			sum -= nums[left]
			left++
		}
	}

	if ret == l+1 {
		return 0
	}

	return ret
}
