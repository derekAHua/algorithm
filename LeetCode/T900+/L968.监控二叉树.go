package T900_

import "math"

// https://leetcode.cn/problems/binary-tree-cameras/

// 题解 https://leetcode.cn/problems/binary-tree-cameras/solutions/422860/jian-kong-er-cha-shu-by-leetcode-solution/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const inf = math.MaxInt64 / 2

func minCameraCover(root *TreeNode) (ret int) {
	var dfs func(*TreeNode) (a, b, c int)

	dfs = func(node *TreeNode) (a, b, c int) {
		if node == nil {
			return inf, 0, 0
		}

		lefta, leftb, leftc := dfs(node.Left)
		righta, rightb, rightc := dfs(node.Right)

		a = leftc + rightc + 1
		b = min(a, min(lefta+rightb, righta+leftb))
		c = min(a, leftb+rightb)

		return
	}

	_, ans, _ := dfs(root)
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
