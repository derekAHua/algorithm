package T000_

//https://leetcode.cn/problems/largest-rectangle-in-histogram/

func largestRectangleArea(heights []int) (max int) {
	stack := make([]int, 0)

	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	stack = append(stack, 0)
	for i, v := range heights {
		for len(stack) > 0 && v < heights[stack[len(stack)-1]] {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				continue
			}

			left := stack[len(stack)-1]
			tmp := heights[mid] * (i - left - 1)
			if tmp > max {
				max = tmp
			}
		}
		stack = append(stack, i)
	}
	return
}
