package T000_

// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/

func searchRange(nums []int, target int) []int {

	leftBorder := getLeft(nums, target)
	rightBorder := getRight(nums, target)

	return []int{leftBorder, rightBorder}
}

func getLeft(nums []int, target int) int {
	l, r := 0, len(nums)
	var mid int

	for l < r {
		mid = l + (r-l)>>1

		if nums[mid] > target {
			r = mid
		} else if nums[mid] == target {
			for mid >= 0 && nums[mid] == target {
				mid--
			}
			if nums[mid+1] != target {
				return -1
			}
			return mid + 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func getRight(nums []int, target int) int {
	l, r := 0, len(nums)
	var mid int

	for l < r {
		mid = l + (r-l)>>1

		if nums[mid] > target {
			r = mid
		} else if nums[mid] == target {
			for mid < len(nums) && nums[mid] == target {
				mid++
			}
			if nums[mid-1] != target {
				return -1
			}
			return mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
