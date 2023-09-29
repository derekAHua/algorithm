package T000_

import "sort"

// https://leetcode.cn/problems/combination-sum-ii/

//力扣题目链接(opens new window)
//
//给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
//candidates 中的每个数字在每个组合中只能使用一次。
//
//说明： 所有数字（包括目标数）都是正整数。解集不能包含重复的组合。
//
//示例 1:
//输入: candidates = [10,1,2,7,6,1,5], target = 8,
//所求解集为:
//[
//  [1, 7],
//  [1, 2, 5],
//  [2, 6],
//  [1, 1, 6]
//]

func combinationSum2(candidates []int, target int) (ret [][]int) {

	sort.Ints(candidates)
	m := make(map[int]int)
	for _, v := range candidates {
		m[v]++
	}

	var f func(idx int)
	var arr []int
	var sum int
	f = func(idx int) {
		if sum == target {
			ret = append(ret, append([]int{}, arr...))
			return
		}
		if sum > target || idx == len(candidates) {
			return
		}

		for i := idx; i < len(candidates); i++ {
			if i > 0 && candidates[i] == candidates[i-1] {
				continue
			}
			v := candidates[i]
			times := m[v]

			for t := 1; t <= times; t++ {
				arr = append(arr, v)
				sum += v
				f(i + 1)
			}
			arr = arr[:len(arr)-times]
			sum -= v * times
		}
	}

	f(0)
	return
}

func combinationSum2R2(candidates []int, target int) (ret [][]int) {
	sort.Ints(candidates)
	var freq [][2]int
	for _, num := range candidates {
		if freq == nil || freq[len(freq)-1][0] != num {
			freq = append(freq, [2]int{num, 1})
		} else {
			freq[len(freq)-1][1]++
		}
	}

	var arr []int
	sum := 0
	var f func(idx int)
	f = func(idx int) {
		if target < sum {
			return
		}
		if target == sum {
			ret = append(ret, append([]int{}, arr...))
			return
		}
		if idx == len(freq) {
			return
		}

		value, time := freq[idx][0], freq[idx][1]
		most := min((target-sum)/value, time)
		for i := 0; i <= most; i++ {
			for j := 0; j < i; j++ {
				arr = append(arr, value)
			}
			sum += value * i
			f(idx + 1)
			sum -= value * i
			arr = arr[:len(arr)-i]
		}
	}

	f(0)
	return
}

func combinationSum2R1(candidates []int, target int) (ret [][]int) {
	sort.Ints(candidates)
	var freq [][2]int
	for _, num := range candidates {
		if freq == nil || freq[len(freq)-1][0] != num {
			freq = append(freq, [2]int{num, 1})
		} else {
			freq[len(freq)-1][1]++
		}
	}

	var arr []int
	var f func(idx, sum int)
	f = func(idx, sum int) {
		if target == sum {
			ret = append(ret, append([]int{}, arr...))
			return
		}
		if idx == len(freq) {
			return
		}
		f(idx+1, sum)
		if target > sum {
			value, time := freq[idx][0], freq[idx][1]
			most := min((target-sum)/value, time)
			for i := 1; i <= most; i++ {
				arr = append(arr, value)
				f(idx+1, sum+value*i)
			}
			arr = arr[:len(arr)-most]
		}
	}
	f(0, 0)
	return
}
