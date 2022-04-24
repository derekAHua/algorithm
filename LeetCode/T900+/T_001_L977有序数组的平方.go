package T900_

// https://leetcode-cn.com/problems/squares-of-a-sorted-array/

func sortedSquares(nums []int) (ret []int) {
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
