package T200_

// https://leetcode.cn/problems/combination-sum-iii/

//找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。
//
//说明：
//
//所有数字都是正整数。
//解集不能包含重复的组合。
//示例 1: 输入: k = 3, n = 7 输出: [[1,2,4]]
//
//示例 2: 输入: k = 3, n = 9 输出: [[1,2,6], [1,3,5], [2,3,4]]

func combinationSum3(k int, n int) (ret [][]int) {
	var arr []int
	var f func(idx, sum int)
	f = func(idx, sum int) {
		if sum > n {
			return
		}
		if len(arr) == k {
			if sum == n {
				ret = append(ret, append([]int{}, arr...))
			}
			return
		}

		for i := idx; i <= 9; i++ {
			arr = append(arr, i)
			f(i+1, sum+i)
			arr = arr[:len(arr)-1]
		}

	}
	f(1, 0)
	return
}

func combinationSum3R1(k int, n int) (ret [][]int) {
	if n < 0 || k < 0 {
		return
	}

	var f func(int, int, []int)

	f = func(start, sum int, arr []int) {
		if sum > n {
			return
		}

		if len(arr) == k && n == sum {
			r := make([]int, len(arr))
			copy(r, arr)
			ret = append(ret, r)
			return
		}

		for i := start; i <= 9; i++ {
			arr = append(arr, i)
			f(i+1, sum+i, arr)
			arr = arr[:len(arr)-1]
		}
	}

	f(1, 0, nil)

	return
}
