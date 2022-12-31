package T700_

// https://leetcode.cn/problems/daily-temperatures/

func dailyTemperatures(temperatures []int) (ret []int) {
	stack := []int{0}
	ret = make([]int, len(temperatures))

	for i := 1; i < len(temperatures); i++ {
		top := stack[len(stack)-1]
		if temperatures[i] <= temperatures[top] {
			stack = append(stack, i)
			continue
		}

		for len(stack) > 0 && temperatures[i] > temperatures[top] {
			ret[top] = i - top
			stack = stack[:len(stack)-1]
			if len(stack) != 0 {
				top = stack[len(stack)-1]
			}
		}
		stack = append(stack, i)
	}

	return
}
