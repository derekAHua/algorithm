package T700_

// https://leetcode-cn.com/problems/binary-search/

func search(nums []int, target int) (ret int) {
	var l, r int
	r = len(nums)
	for l != r {
		mid := l + (-l+r)/2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			l = mid + 1
		default:
			r = mid
		}
	}

	return -1
}

func searchRepeat4(nums []int, target int) (ret int) {

	l := 0
	r := len(nums)
	for l < r {
		cur := (l + r) >> 1
		if nums[cur] == target {
			return cur
		}
		if nums[cur] > target {
			r = cur
			continue
		}
		l = cur + 1
	}

	return -1
}

func searchRepeat3(nums []int, target int) (ret int) {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)>>1

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid
			continue
		}

		left = mid + 1
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

func searchV1(nums []int, target int) (ret int) {
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

func searchRepeat2(nums []int, target int) (ret int) {

	var left, right, mid, midNumber = 0, len(nums), 0, 0

	//[)
	for left < right {
		mid = left + (right-left)>>1
		midNumber = nums[mid]

		if midNumber == target {
			return mid
		}

		if midNumber > target {
			right = mid
			continue
		}

		left = mid + 1
	}

	return -1
}
