package T900_

import "sort"

// https://leetcode-cn.com/problems/squares-of-a-sorted-array/

func sortedSquares(nums []int) (ret []int) {
	ret = make([]int, len(nums))
	l, r := 0, len(nums)-1
	cur := len(nums) - 1
	for l <= r {
		lValue, rValue := nums[l]*nums[l], nums[r]*nums[r]
		if lValue > rValue {
			l++
			ret[cur] = lValue
			cur--
			continue
		}
		ret[cur] = rValue
		cur--
		r--
	}

	return
}

func sortedSquaresRepeat4(nums []int) (ret []int) {

	for i, v := range nums {
		if v < 0 {
			nums[i] = (-v) * (-v)
			continue
		}
		nums[i] = v * v
	}

	sort.Ints(nums)

	return nums
}

func sortedSquaresRepeat3(nums []int) (ret []int) {
	length := len(nums)
	ret = make([]int, length)

	l, r, cur := 0, length-1, length-1
	for ; l <= r; cur-- {
		lValue, rValue := nums[l]*nums[l], nums[r]*nums[r]
		if lValue > rValue {
			ret[cur] = lValue
			l++
			continue
		}
		ret[cur] = rValue
		r--
	}

	return
}

func sortedSquaresRepeat2(nums []int) (ret []int) {
	l := len(nums)
	ret = make([]int, l)

	left, right, cur := 0, l-1, l-1
	for left <= right {
		lm, rm := nums[left]*nums[left], nums[right]*nums[right]
		if lm >= rm {
			ret[cur] = lm
			left++
		} else {
			ret[cur] = rm
			right--
		}
		cur--
	}

	return ret
}

func sortedSquaresRepeat1(nums []int) (ret []int) {

	maxIndex := len(nums) - 1

	l, r := 0, maxIndex

	ret = make([]int, len(nums))

	for ; l != r; maxIndex-- {
		leftV := nums[l] * nums[l]
		rightV := nums[r] * nums[r]

		if rightV >= leftV {
			ret[maxIndex] = rightV
			r--
			continue
		}

		ret[maxIndex] = leftV
		l++
	}

	ret[maxIndex] = nums[l] * nums[l]
	return
}
