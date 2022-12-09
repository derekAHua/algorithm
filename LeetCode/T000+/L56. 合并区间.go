package T000_

import "sort"

// https://leetcode.cn/problems/merge-intervals/

func merge(intervals [][]int) (ret [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][1] >= intervals[i+1][0] {
			intervals[i+1] = []int{mi(intervals[i][0], intervals[i+1][0]), m(intervals[i][1], intervals[i+1][1])}
			continue
		}

		ret = append(ret, intervals[i])
	}
	ret = append(ret, intervals[len(intervals)-1])

	return
}

func m(x, y int) int {
	if y > x {
		return y
	}
	return x
}

func mi(x, y int) int {
	if y > x {
		return x
	}
	return y
}
