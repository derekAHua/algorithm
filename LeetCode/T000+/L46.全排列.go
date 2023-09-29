package T000_

// https://leetcode.cn/problems/permutations/

//给定一个 没有重复 数字的序列，返回其所有可能的全排列。
//
//示例:
//
//输入: [1,2,3]
//输出: [ [1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1] ]

func permute(nums []int) (ret [][]int) {

	var (
		f   func()
		arr []int
	)

	m := map[int]bool{}
	f = func() {
		if len(arr) == len(nums) {
			ret = append(ret, append([]int{}, arr...))
		}

		for i := 0; i < len(nums); i++ {
			if m[nums[i]] {
				continue
			}
			arr = append(arr, nums[i])
			m[nums[i]] = true
			f()
			arr = arr[:len(arr)-1]
			m[nums[i]] = false
		}
	}

	f()
	return
}

func permuteR2(nums []int) (ret [][]int) {
	m := map[int]bool{}
	var f func()
	var arr []int
	f = func() {
		if len(arr) == len(nums) {
			ret = append(ret, append([]int{}, arr...))
		}

		for i := 0; i < len(nums); i++ {
			if m[i] {
				continue
			}
			arr = append(arr, nums[i])
			m[i] = true
			f()
			m[i] = false
			arr = arr[:len(arr)-1]
		}
	}

	f()
	return
}

func permuteR1(nums []int) (ret [][]int) {
	var f func(map[int]bool, []int)
	f = func(m map[int]bool, arr []int) {
		if len(arr) == len(nums) {
			r := make([]int, len(arr))
			copy(r, arr)
			ret = append(ret, r)
		}

		for i := 0; i < len(nums); i++ {
			if m[nums[i]] {
				continue
			}
			arr = append(arr, nums[i])
			m[nums[i]] = true
			f(m, arr)
			m[nums[i]] = false
			arr = arr[:len(arr)-1]
		}
	}

	f(map[int]bool{}, nil)
	return
}
