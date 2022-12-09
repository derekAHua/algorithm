package T400_

import (
	"sort"
)

// https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/

func findMinArrowShots(points [][]int) (ret int) {
	ret = 1
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	for i := 1; i < len(points); i++ {
		if points[i-1][1] < points[i][0] {
			ret++
			continue
		}

		points[i][1] = min(points[i-1][1], points[i][1])
	}

	return
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
