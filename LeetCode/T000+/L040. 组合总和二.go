package T000_

import "sort"

// https://leetcode.cn/problems/combination-sum-ii/

func combinationSum2(candidates []int, target int) (ret [][]int) {
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
