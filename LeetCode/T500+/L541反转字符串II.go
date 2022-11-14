package T500_

// https://leetcode.cn/problems/reverse-string-ii/

//给定一个字符串 s 和一个整数 k，从字符串开头算起, 每计数至 2k 个字符，就反转这 2k 个字符中的前 k 个字符。
//
//如果剩余字符少于 k 个，则将剩余字符全部反转。
//
//如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
//
//示例:
//
//输入: s = "abcdefg", k = 2
//输出: "bacdfeg"

func reverseStr(s string, k int) string {
	str := []byte(s)
	length := len(s)
	for i := 0; i < length; i += 2 * k {
		r1 := i + k

		if length < r1 {
			reverseString(str[i:length])
			continue
		}

		reverseString(str[i:r1])
	}

	return string(str)
}

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
