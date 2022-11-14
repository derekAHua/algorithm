package T200_

// https://leetcode.cn/problems/4sum-ii/

//给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组 (i, j, k, l) ，使得 A[i] + B[j] + C[k] + D[l] = 0。
//
//为了使问题简单化，所有的 A, B, C, D 具有相同的长度 N，且 0 ≤ N ≤ 500 。
//所有整数的范围在 -2^28 到 2^28 - 1 之间，最终结果不会超过 2^31 - 1 。

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int)
	ret := 0
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			m[v1+v2]++
		}
	}
	for _, v1 := range nums3 {
		for _, v2 := range nums4 {
			ret += m[-v1-v2]
		}
	}

	return ret
}
