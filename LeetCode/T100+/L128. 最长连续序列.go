package T100_

import "sort"

func longestConsecutive(nums []int) (ret int) {
	m := make(map[int]bool, len(nums)/2)
	for _, v := range nums {
		m[v] = true
	}

	for key := range m {
		if !m[key-1] {
			cnt := 1
			for m[key+1] {
				cnt++
				key++
			}
			if cnt > ret {
				ret = cnt
			}
		}
	}

	return
}

func longestConsecutiveR2(nums []int) (ret int) {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	cnt := 1
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 0 {
			continue
		}
		if nums[i]-nums[i-1] == 1 {
			cnt++
		} else {
			if ret < cnt {
				ret = cnt
			}
			cnt = 1
		}
	}

	if ret < cnt {
		ret = cnt
	}

	return
}
