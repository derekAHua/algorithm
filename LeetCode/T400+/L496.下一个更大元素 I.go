package T400_

// https://leetcode.cn/problems/next-greater-element-i/

func nextGreaterElement(nums1 []int, nums2 []int) (ret []int) {
	ret = make([]int, len(nums1))
	for i := range ret {
		ret[i] = -1
	}

	m := map[int]int{}
	for i, v := range nums2 {
		m[v] = i
	}

	for i, v := range nums1 {
		for j := m[v]; j < len(nums2); j++ {
			if nums2[j] > nums1[i] {
				ret[i] = nums2[j]
				break
			}
		}
	}

	return
}

func nextGreaterElementR2(nums1 []int, nums2 []int) (ret []int) {
	ret = make([]int, 0, len(nums1))

	for _, v := range nums1 {
		f := false
		yes := false
		for i := 0; i < len(nums2); i++ {
			if nums2[i] == v {
				f = true
				continue
			}
			if f && nums2[i] > v {
				ret = append(ret, nums2[i])
				yes = true
				break
			}
		}
		if !yes {
			ret = append(ret, -1)
		}
	}
	return
}

func nextGreaterElementR1(nums1 []int, nums2 []int) (ret []int) {
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
