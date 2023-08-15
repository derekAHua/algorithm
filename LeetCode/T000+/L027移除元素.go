package T000_

// https://leetcode-cn.com/problems/remove-element/

func removeElement(nums []int, val int) (ret int) {
	idx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[idx] = nums[i]
			idx++
		}
	}

	return idx + 1
}

func removeElementRepeat3(nums []int, val int) (ret int) {
	if len(nums) == 0 {
		return
	}

	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}
		ret++
		nums[j] = nums[i]
		j++
	}
	return
}

func removeElement2(nums []int, val int) int {
	l := len(nums)
	slow := 0                // 慢指针
	for i := 0; i < l; i++ { // i:快指针
		if nums[i] != val {
			nums[slow] = nums[i]
			slow++
		}
	}

	return slow
}

func removeElementRepeat1(nums []int, val int) int {

	var slowIndex, fastIndex int

	for fastIndex < len(nums) {
		if nums[fastIndex] == val {
			fastIndex++
			continue
		}

		nums[slowIndex] = nums[fastIndex]
		fastIndex++
		slowIndex++
	}

	return slowIndex
}

func removeElementRepeat2(nums []int, val int) int {
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == val {
			continue
		}
		nums[left] = nums[right]
		left++
	}

	nums = nums[:left]

	return left
}
