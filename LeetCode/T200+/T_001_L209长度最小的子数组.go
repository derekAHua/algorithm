package T200_

// https://leetcode-cn.com/problems/minimum-size-subarray-sum/

func minSubArrayLen(target int, nums []int) (ret int) {
	l := len(nums)
	if l == 0 {
		return 0
	}

	left, sum := 0, 0        // left：滑动窗口左指针
	ret = l + 1              // 初始化返回长度为l+1，目的是为了判断“不存在符合条件的子数组，返回0”的情况
	for i := 0; i < l; i++ { // i：滑动窗口右指针
		sum += nums[i]
		for sum >= target {
			subLength := i - left + 1
			if subLength < ret {
				ret = subLength
			}
			sum -= nums[left]
			left++
		}
	}

	if ret == l+1 {
		return 0
	}

	return ret
}

func minSubArrayLenRepeat1(target int, nums []int) (ret int) {
	length := len(nums)
	if length == 0 {
		return 0
	}

	l, r := 0, 0

	sum := nums[0]
	ret = length + 1

	for r < length {
		for sum < target {
			r++
			if r >= length {
				if ret == length+1 {
					return 0
				}
				return
			}
			sum += nums[r]
		}

		ret = getMin(ret, r-l+1)

		for sum >= target {
			sum -= nums[l]
			l++
		}

		ret = getMin(ret, r-l+1+1)
	}

	return
}

func getMin(x, y int) int {
	if x > y {
		return y
	}
	return x
}
