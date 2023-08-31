package T500_

// https://leetcode.cn/problems/next-greater-element-ii/

func nextGreaterElements(nums []int) (ret []int) {
	ret = make([]int, len(nums))
	for i := range ret {
		ret[i] = -1
	}

	var stack []int
	nums = append(nums, nums...)
	for i, v := range nums {
		for len(stack) > 0 && v > nums[stack[len(stack)-1]%len(ret)] {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret[pop%len(ret)] = v
		}

		stack = append(stack, i)
	}

	return
}

func nextGreaterElementsR3(nums []int) (ret []int) {

	var stack []int
	ret = make([]int, len(nums))
	for i := range ret {
		ret[i] = -1
	}

	for i := 0; i < len(nums)*2; i++ {
		v := nums[i%len(nums)]
		for len(stack) > 0 && v > nums[stack[len(stack)-1]] {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret[pop] = v
		}

		stack = append(stack, i%len(nums))
	}

	return
}

func nextGreaterElementsR2(nums []int) (ret []int) {
	length := len(nums)
	ret = make([]int, length)
	for i := range ret {
		ret[i] = -1
	}

	nums = append(nums, nums...)
	var stack []int
	for i := 0; i < length*2; i++ {
		v := nums[i%length]
		for len(stack) > 0 && v > nums[stack[len(stack)-1]] {
			ret[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%length)
	}

	return
}

func nextGreaterElementsR1(nums []int) (ret []int) {
	length := len(nums)
	ret = make([]int, length)
	for i := range ret {
		ret[i] = -1
	}

	nums = append(nums, nums...)
	var stack []int
	for i, v := range nums {
		for len(stack) > 0 && v > nums[stack[len(stack)-1]] {
			ret[stack[len(stack)-1]%length] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return
}
