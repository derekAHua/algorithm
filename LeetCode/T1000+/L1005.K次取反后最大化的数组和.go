package T1000_

import (
	"math"
	"sort"
)

// https://leetcode.cn/problems/maximize-sum-of-array-after-k-negations/

// 给定一个整数数组 A，我们只能用以下方法修改该数组：我们选择某个索引 i 并将 A[i] 替换为 -A[i]，然后总共重复这个过程 K 次。（我们可以多次选择同一个索引 i。）
//
//以这种方式修改数组后，返回数组可能的最大和。
//
//示例 1：
//
//输入：A = [4,2,3], K = 1
//输出：5
//解释：选择索引 (1,) ，然后 A 变为 [4,-2,3]。
//示例 2：
//
//输入：A = [3,-1,0,2], K = 3
//输出：6
//解释：选择索引 (1, 2, 2) ，然后 A 变为 [3,1,0,2]。

func largestSumAfterKNegations(nums []int, K int) int {
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})

	for i := 0; i < len(nums); i++ {
		if K > 0 && nums[i] < 0 {
			nums[i] = -nums[i]
			K--
		}
	}

	if K%2 == 1 {
		nums[len(nums)-1] = -nums[len(nums)-1]
	}

	result := 0
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return result
}
