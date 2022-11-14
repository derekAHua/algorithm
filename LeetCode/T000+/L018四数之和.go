package T000_

import "sort"

// https://leetcode.cn/problems/4sum/

//题意：给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，
//使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
//
//注意：
//
//答案中不可以包含重复的四元组。
//
//示例： 给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。 满足要求的四元组集合为： [ [-1, 0, 0, 1], [-2, -1, 1, 2], [-2, 0, 0, 2] ]

func fourSum(nums []int, target int) [][]int {
	length := len(nums)
	if length < 4 {
		return nil
	}
	sort.Ints(nums)
	ret := make([][]int, 0)

	for i := 0; i < length-3; i++ {
		v1 := nums[i]
		if i > 0 && nums[i-1] == v1 {
			continue
		}

		for j := i + 1; j < length-2; j++ {
			v2 := nums[j]
			if j > i+1 && nums[j-1] == v2 {
				continue
			}

			left := j + 1
			right := length - 1
			for left < right {
				v3 := nums[left]
				v4 := nums[right]

				sum := v1 + v2 + v3 + v4
				if sum > target {
					right--
					continue
				}
				if sum < target {
					left++
					continue
				}
				ret = append(ret, []int{v1, v2, v3, v4})
				for left < right && nums[left] == v3 {
					left++
				}
				for left < right && nums[right] == v4 {
					right--
				}
			}
		}
	}

	return ret
}
