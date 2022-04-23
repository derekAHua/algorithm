package T700_

// https://leetcode-cn.com/problems/binary-search/

// Do not like.
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

func searchRepeat1(nums []int, target int) (ret int) {
	ret = -1

	l, r := 0, len(nums)

	for l < r {
		mid := l + (r-l)>>1

		if nums[mid] > target {
			r = mid
		} else if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		}
	}

	return
}
