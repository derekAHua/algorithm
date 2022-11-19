package T100_

// https://leetcode.cn/problems/path-sum-ii/

func pathSum(root *TreeNode, targetSum int) (ret [][]int) {
	var f func(node *TreeNode, arr []int, targetSum int)

	f = func(node *TreeNode, arr []int, target int) {
		if node == nil {
			return
		}

		ints := append(arr, node.Val)
		target -= node.Val
		if node.Left == nil && node.Right == nil && target == 0 {
			r := make([]int, 0)
			for _, cv := range ints {
				r = append(r, cv)
			}
			ret = append(ret, r)
			return
		}

		if node.Left != nil {
			f(node.Left, ints, target)
		}
		if node.Right != nil {
			f(node.Right, ints, target)
		}
		ints = ints[:len(ints)-1]
	}

	f(root, nil, targetSum)

	return
}
