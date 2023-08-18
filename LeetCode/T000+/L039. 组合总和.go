package T000_

// https://leetcode.cn/problems/combination-sum/

func combinationSum(candidates []int, target int) (ret [][]int) {
	if len(candidates) == 0 {
		return
	}

	sum := 0
	var arr []int
	var f func(idx int)

	f = func(idx int) {
		if sum > target {
			return
		}
		if sum == target {
			ret = append(ret, append([]int{}, arr...))
			return
		}
		if idx == len(candidates) {
			return
		}

		value := candidates[idx]
		for i := 0; i <= target/value; i++ {
			sum += value * i
			for j := 0; j < i; j++ {
				arr = append(arr, value)
			}
			f(idx + 1)
			arr = arr[:len(arr)-i]
			sum -= value * i
		}
	}

	f(0)
	return
}
