package T000_

// https://leetcode.cn/problems/combinations/

//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
//
//示例:
//输入: n = 4, k = 2
//输出:
//[
//[2,4],
//[3,4],
//[2,3],
//[1,2],
//[1,3],
//[1,4],
//]
//

func combine(n int, k int) (ret [][]int) {
	var arr []int
	var f func(idx int)
	f = func(idx int) {
		if len(arr) == k {
			ret = append(ret, append([]int{}, arr...))
			return
		}
		if idx == n+1 {
			return
		}

		for i := idx; i <= n; i++ {
			arr = append(arr, i)
			f(i + 1)
			arr = arr[:len(arr)-1]
		}
	}
	f(1)
	return
}

func combineR2(n int, k int) (ret [][]int) {
	var arr []int
	var f func(idx int)
	f = func(idx int) {
		if idx == n+1 {
			if len(arr) == k {
				ret = append(ret, append([]int{}, arr...))
			}
			return
		}

		f(idx + 1)
		arr = append(arr, idx)
		f(idx + 1)
		arr = arr[:len(arr)-1]
	}
	f(1)
	return
}

func combineR1(n int, k int) (ret [][]int) {
	if n < 1 || k > n {
		return
	}

	var f func(int, []int)

	f = func(start int, arr []int) {
		if len(arr) == k {
			t := make([]int, len(arr))
			copy(t, arr)
			ret = append(ret, t)
			return
		}

		if len(arr)+n-start+1 < k {
			return
		}

		for i := start; i <= n; i++ {
			arr = append(arr, i)
			f(i+1, arr)
			arr = arr[:len(arr)-1]
		}
	}

	f(1, nil)

	return
}
