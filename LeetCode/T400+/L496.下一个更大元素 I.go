package T400_

// https://leetcode.cn/problems/next-greater-element-i/

func nextGreaterElement(nums1 []int, nums2 []int) (ret []int) {
	ret = make([]int, len(nums1))
	m := make(map[int]int)

	stack := []int{0}
	for i := 1; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[i] > nums2[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			m[nums2[top]] = nums2[i]

			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	for i, v := range nums1 {
		if num, ok := m[v]; ok {
			ret[i] = num
		} else {
			ret[i] = -1
		}
	}

	return
}
