package T000_

// @Author: Derek
// @Description:
// @Date: 2023/8/15 23:25
// @Version 1.0

// https://leetcode.cn/problems/trapping-rain-water/

func trap(height []int) (ret int) {
	var stack []int

	for i, v := range height {
		for len(stack) > 0 && v >= height[stack[len(stack)-1]] {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				break
			}

			lIndex := stack[len(stack)-1]
			ret += (min(v, height[lIndex]) - height[pop]) * (i - lIndex - 1)
		}

		stack = append(stack, i)
	}

	return
}

func trapR1(height []int) (ret int) {

	var stack []int
	for i := 0; i < len(height); i++ {
		v := height[i]

		for len(stack) > 0 && v >= height[stack[len(stack)-1]] {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				break
			}

			lIndex := stack[len(stack)-1]
			ret += (min(v, height[lIndex]) - height[pop]) * (i - lIndex - 1)
		}

		stack = append(stack, i)
	}

	return
}
