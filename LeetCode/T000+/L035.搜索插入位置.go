package T000_

// https://leetcode-cn.com/problems/search-insert-position/

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)
	var mid int

	for l < r {
		mid = l + (r-l)>>1

		if nums[mid] > target {
			r = mid
		} else if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		}
	}

	return r
}
