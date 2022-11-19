package T500_

import "sort"

// https://leetcode.cn/problems/find-mode-in-binary-search-tree/

func findMode(root *TreeNode) (ret []int) {
	if root == nil {
		return
	}

	m := make(map[int]int)
	list := make([]int, 0)
	var f func(node *TreeNode)
	f = func(node *TreeNode) {
		if node == nil {
			return
		}
		m[node.Val]++
		list = append(list, node.Val)
		f(node.Left)
		f(node.Right)
	}

	f(root)

	sort.Slice(list, func(i, j int) bool {
		return m[list[i]] > m[list[j]]
	})

	value := m[list[0]]
	for k, v := range m {
		if v == value {
			ret = append(ret, k)
		}
	}

	return
}
