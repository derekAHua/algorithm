package T500_

import (
	"math"
	"sort"
)

// https://leetcode.cn/problems/minimum-absolute-difference-in-bst/

func getMinimumDifference(root *TreeNode) (ret int) {
	ret = math.MaxInt

	var list []int
	var f func(node *TreeNode)
	f = func(node *TreeNode) {
		if node == nil {
			return
		}
		list = append(list, node.Val)
		f(node.Left)
		f(node.Right)
	}

	f(root)

	sort.Ints(list)

	for i := 1; i < len(list); i++ {
		v := list[i] - list[i-1]
		if v < ret {
			ret = v
		}
	}

	return
}
