package T400_

import (
	"sort"
)

// https://leetcode.cn/problems/non-overlapping-intervals/

func eraseOverlapIntervals(intervals [][]int) (ret int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			ret++
			intervals[i][1] = mi(intervals[i-1][1], intervals[i][1])
		}
	}

	return
}

func mi(a, b int) int {
	if a > b {
		return b
	}
	return a
}
