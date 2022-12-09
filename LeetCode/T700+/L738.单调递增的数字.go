package T700_

import "strconv"

// https://leetcode.cn/problems/monotone-increasing-digits/

func monotoneIncreasingDigits(n int) (ret int) {
	s := strconv.Itoa(n)
	arr := []byte(s)
	l := len(arr)
	if l <= 1 {
		return n
	}

	for i := l - 1; i > 0; i-- {
		if arr[i] < arr[i-1] {
			arr[i-1]--
			for j := i; j < l; j++ {
				arr[j] = '9'
			}
		}
	}

	ret, _ = strconv.Atoi(string(arr))

	return
}
