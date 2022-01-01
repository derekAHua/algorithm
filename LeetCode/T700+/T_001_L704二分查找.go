package T700_

// https://leetcode-cn.com/problems/binary-search/

func searchV1(nums []int, target int) int {
	for index, v := range nums {
		if v == target {
			return index
		}
	}
	return -1
}

func searchV2(nums []int, target int) int {
	high := len(nums) - 1
	low := 0
	// 区间：[left,right]
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func search(nums []int, target int) (ret int) {
	right := len(nums)
	left := 0

	var mid int

	// 区间：[left,right)
	for left < right {
		mid = left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return -1
}
