package T000_

import "sort"

// https://leetcode.cn/problems/subsets-ii/

//给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
//
//说明：解集不能包含重复的子集。
//
//示例:
//
//输入: [1,2,2]
//输出: [ [2], [1], [1,2,2], [2,2], [1,2], [] ]

func subsetsWithDup(nums []int) (ret [][]int) {

	sort.Ints(nums)
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	var f func(idx int)
	var arr []int

	f = func(idx int) {
		ret = append(ret, append([]int{}, arr...))

		for i := idx; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			v := nums[i]
			times := m[v]

			for t := 1; t <= times; t++ {
				arr = append(arr, v)
				f(i + 1)
			}
			arr = arr[:len(arr)-times]
		}
	}

	f(0)
	return
}

func subsetsWithDupR2(nums []int) (ret [][]int) {
	sort.Ints(nums)
	var arr []int
	var f func(idx int, choosePre bool)
	f = func(idx int, choosePre bool) {
		if idx == len(nums) {
			ret = append(ret, append([]int{}, arr...))
			return
		}

		f(idx+1, false)

		if !choosePre && idx > 0 && nums[idx] == nums[idx-1] {
			return
		}
		arr = append(arr, nums[idx])
		f(idx+1, true)
		arr = arr[:len(arr)-1]
	}

	f(0, false)
	return
}

func subsetsWithDupR1(nums []int) (ret [][]int) {
	sort.Ints(nums)

	var f func(int, []int)

	f = func(index int, arr []int) {
		r := make([]int, len(arr))
		copy(r, arr)
		ret = append(ret, r)

		for i := index; i < len(nums); i++ {
			if i > index && nums[i] == nums[i-1] {
				continue
			}

			arr = append(arr, nums[i])
			f(i+1, arr)
			arr = arr[:len(arr)-1]
		}
	}

	f(0, nil)
	return
}
