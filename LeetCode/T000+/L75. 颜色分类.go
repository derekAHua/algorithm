package T000_

func sortColors(nums []int) {

	cur := 0
	idx := 0
	for i := idx; i < len(nums); i++ {
		if nums[i] == cur {
			nums[i], nums[idx] = nums[idx], nums[i]
			idx++
		}
	}

	cur = 1
	for i := idx; i < len(nums); i++ {
		if nums[i] == cur {
			nums[i], nums[idx] = nums[idx], nums[i]
			idx++
		}
	}

	return
}
