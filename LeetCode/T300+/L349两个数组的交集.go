package T300_

// https://leetcode.cn/problems/intersection-of-two-arrays/

// 题意：给定两个数组，编写一个函数来计算它们的交集。

func intersection(nums1 []int, nums2 []int) []int {
	ret := make([]int, 0)

	m := make(map[int]struct{}, len(nums1))
	for _, v := range nums1 {
		m[v] = struct{}{}
	}

	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			ret = append(ret, v)
			delete(m, v)
		}
	}

	return ret
}
