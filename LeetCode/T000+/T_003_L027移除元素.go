package T000_

// https://leetcode-cn.com/problems/remove-element/

func removeElement(nums []int, val int) int {
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
