package T400_

// https://leetcode.cn/problems/increasing-subsequences/

// 给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是2。
//
// 示例:
//
// 输入: [4, 6, 7, 7]
// 输出: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7]]
// 说明:
//
// 给定数组的长度不会超过15。
// 数组中的整数范围是 [-100,100]。
// 给定数组中可能包含重复数字，相等的数字应该被视为递增的一种情况。
func findSubsequences(nums []int) (ret [][]int) {
	var f func(idx int)
	var arr []int

	f = func(idx int) {
		if len(arr) >= 2 {
			ret = append(ret, append([]int{}, arr...))
		}

		history := [201]bool{}
		for i := idx; i < len(nums); i++ {
			if len(arr) > 0 && nums[i] < arr[len(arr)-1] || history[nums[i]+100] {
				continue
			}
			history[nums[i]+100] = true

			arr = append(arr, nums[i])
			f(i + 1)
			arr = arr[:len(arr)-1]
		}
	}

	f(0)
	return
}

func findSubsequencesR1(nums []int) (ret [][]int) {
	var f func(int, []int)
	f = func(index int, arr []int) {
		if len(arr) >= 2 {
			r := make([]int, len(arr))
			copy(r, arr)
			ret = append(ret, r)
		}

		history := [201]bool{}
		for i := index; i < len(nums); i++ {
			//分两种情况判断：一，当前取的元素小于子集的最后一个元素，则继续寻找下一个适合的元素
			//                或者二，当前取的元素在本层已经出现过了，所以跳过该元素，继续寻找
			if len(arr) > 0 && nums[i] < arr[len(arr)-1] || history[nums[i]+100] == true {
				continue
			}
			history[nums[i]+100] = true
			arr = append(arr, nums[i])
			f(i+1, arr)
			arr = arr[:len(arr)-1]
		}
	}

	f(0, nil)
	return
}
