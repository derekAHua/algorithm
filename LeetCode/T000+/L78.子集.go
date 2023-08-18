package T000_

// https://leetcode.cn/problems/subsets/

//给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
//
//说明：解集不能包含重复的子集。
//
//示例: 输入: nums = [1,2,3] 输出: [ [3],   [1],   [2],   [1,2,3],   [1,3],   [2,3],   [1,2],   [] ]

func subsets(nums []int) (ret [][]int) {
	var arr []int
	var f func(idx int)

	f = func(idx int) {
		if idx == len(nums) {
			ret = append(ret, append([]int{}, arr...))
			return
		}

		f(idx + 1)
		arr = append(arr, nums[idx])
		f(idx + 1)
		arr = arr[:len(arr)-1]
	}
	f(0)
	return
}

func subsets2(nums []int) (ret [][]int) {

	var f func(index int, arr []int)
	f = func(index int, arr []int) {
		r := make([]int, len(arr))
		copy(r, arr)
		ret = append(ret, r)

		for i := index; i < len(nums); i++ {
			arr = append(arr, nums[i])
			f(i+1, arr)
			arr = arr[:len(arr)-1]
		}
	}

	f(0, nil)

	return
}
