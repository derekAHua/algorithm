package T700_

// @Author: Derek
// @Description:
// @Date: 2023/8/27 21:12
// @Version 1.0

func findLength(nums1 []int, nums2 []int) (ret int) {
	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > ret {
				ret = dp[i][j]
			}
		}
	}

	return
}

func findLengthR2(nums1 []int, nums2 []int) (ret int) {
	dp := make([][]int, len(nums1))
	for i := range dp {
		dp[i] = make([]int, len(nums2))
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if nums1[i] == nums2[j] {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j-1] + 1
				}
			}
			if dp[i][j] > ret {
				ret = dp[i][j]
			}
		}
	}

	return
}

func findLengthR1(nums1 []int, nums2 []int) (ret int) {
	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > ret {
				ret = dp[i][j]
			}
		}
	}

	return
}
