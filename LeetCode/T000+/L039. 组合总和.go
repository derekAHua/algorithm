package T000_

// https://leetcode.cn/problems/combination-sum/

//给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
//candidates 中的数字可以无限制重复被选取。
//
//说明：
//
//所有数字（包括 target）都是正整数。
//解集不能包含重复的组合。
//示例 1：
//
//输入：candidates = [2,3,6,7], target = 7,
//所求解集为： [ [7], [2,2,3] ]
//示例 2：
//
//输入：candidates = [2,3,5], target = 8,
//所求解集为： [ [2,2,2,2], [2,3,3], [3,5] ]
//#

func combinationSum(candidates []int, target int) (ret [][]int) {
	var (
		arr []int
		sum int
		f   func(idx int)
	)
	f = func(idx int) {
		if sum == target {
			ret = append(ret, append([]int{}, arr...))
			return
		}
		if sum > target || idx == len(candidates) {
			return
		}
		for i := idx; i < len(candidates); i++ {
			times := (target - sum) / candidates[i]
			if times == 0 {
				continue
			}
			for j := 0; j < times; j++ {
				arr = append(arr, candidates[i])
				sum += candidates[i]
				f(i + 1)
			}
			arr = arr[:len(arr)-times]
			sum -= times * candidates[i]
		}
	}
	f(0)
	return
}

func combinationSumR3(candidates []int, target int) (ret [][]int) {

	var f func(idx int)
	sum := 0
	var arr []int
	f = func(idx int) {
		if sum == target {
			ret = append(ret, append([]int{}, arr...))
			return
		}

		if sum > target || idx == len(candidates) {
			return
		}

		for i := idx; i < len(candidates); i++ {
			arr = append(arr, candidates[i])
			sum += candidates[i]
			f(i)
			arr = arr[:len(arr)-1]
			sum -= candidates[i]
		}
	}

	f(0)
	return
}

func combinationSumR2(candidates []int, target int) (ret [][]int) {

	var arr []int
	sum := 0
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

		times := (target - sum) / candidates[idx]
		for i := 0; i <= times; i++ {
			for j := 0; j < i; j++ {
				arr = append(arr, candidates[idx])
			}
			sum += candidates[idx] * i
			f(idx + 1)
			sum -= candidates[idx] * i
			arr = arr[:len(arr)-i]
		}
	}

	f(0)
	return
}

func combinationSumR1(candidates []int, target int) (ret [][]int) {
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
