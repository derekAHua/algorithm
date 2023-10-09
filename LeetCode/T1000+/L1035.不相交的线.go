package T1000_

// https://leetcode.cn/problems/uncrossed-lines/

func maxUncrossedLines(nums1 []int, nums2 []int) (ret int) {
	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}

func maxUncrossedLinesR2(nums1 []int, nums2 []int) (ret int) {
	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
			if ret < dp[i][j] {
				ret = dp[i][j]
			}
		}
	}

	return
}

func maxUncrossedLinesR1(nums1 []int, nums2 []int) int {

	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(nums1)][len(nums2)]
}
