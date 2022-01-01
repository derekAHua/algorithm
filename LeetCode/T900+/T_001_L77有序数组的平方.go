package T900_

import (
	"sort"
)

// https://leetcode-cn.com/problems/squares-of-a-sorted-array/

/*
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

示例 1： 输入：nums = [-4,-1,0,3,10] 输出：[0,1,9,16,100] 解释：平方后，数组变为 [16,1,0,9,100]，排序后，数组变为 [0,1,9,16,100]

示例 2： 输入：nums = [-7,-3,2,3,11] 输出：[4,9,9,49,121]
*/

func sortedSquaresV1(nums []int) []int {
	for index, v := range nums {
		nums[index] *= v
	}
	sort.Ints(nums)
	return nums
}

func sortedSquares(nums []int) (ret []int) {
	l := len(nums)
	ret = make([]int, l)

	left, right, cur := 0, l-1, l-1
	for left <= right {
		lm, rm := nums[left]*nums[left], nums[right]*nums[right]
		if lm >= rm {
			ret[cur] = lm
			left++
		} else {
			ret[cur] = rm
			right--
		}
		cur--
	}

	return ret
}
