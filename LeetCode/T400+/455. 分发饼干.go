package T400_

import "sort"

// https://leetcode.cn/problems/assign-cookies/

func findContentChildren(g []int, s []int) (ret int) {
	sort.Ints(g)
	sort.Ints(s)

	index := 0
	for _, v := range g {
		for index < len(s) && v > s[index] {
			index++
		}
		if index >= len(s) {
			return
		}

		index++
		ret++
	}

	return
}
