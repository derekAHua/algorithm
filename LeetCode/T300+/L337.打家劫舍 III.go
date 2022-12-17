package T300_

// https://leetcode.cn/problems/house-robber-iii/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {

	var f func(node *TreeNode) []int

	f = func(node *TreeNode) (ret []int) {
		if node == nil {
			return []int{0, 0}
		}

		left := f(node.Left)
		right := f(node.Right)

		robCur := node.Val + left[0] + right[0]
		notRobCur := max(left[0], left[1]) + max(right[0], right[1])
		return []int{notRobCur, robCur}
	}

	ints := f(root)
	return max(ints[0], ints[1])
}
