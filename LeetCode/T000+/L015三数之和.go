package T000_

import "sort"

// https://leetcode.cn/problems/3sum/

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
//
//注意： 答案中不可以包含重复的三元组。
//
//示例：
//
//给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为： [ [-1, 0, 1], [-1, -1, 2] ]

func threeSum(nums []int) [][]int {
	length := len(nums)
	if length < 3 {
		return nil
	}
	sort.Ints(nums)

	ret := make([][]int, 0)

	for i := 0; i < length-2; i++ {
		iValue := nums[i]
		if iValue > 0 {
			break
		}
		if i > 0 && iValue == nums[i-1] {
			continue
		}

		left := i + 1
		right := length - 1
		for left < right {
			leftV := nums[left]
			rightV := nums[right]
			sum := iValue + leftV + rightV
			if sum > 0 {
				right--
				continue
			}
			if sum < 0 {
				left++
				continue
			}
			ret = append(ret, []int{iValue, leftV, rightV})
			for left < right && nums[left] == leftV {
				left++
			}
			for left < right && nums[right] == rightV {
				right--
			}
		}
	}

	return ret
}
