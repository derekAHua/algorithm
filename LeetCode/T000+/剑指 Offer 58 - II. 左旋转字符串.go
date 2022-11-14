package T000_

// @Author: Derek
// @Description:
// @Date: 2022/11/13 14:59
// @Version 1.0

// https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/

//字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。
//请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。
//
//示例 1：
//输入: s = "abcdefg", k = 2
//输出: "cdefgab"
//
//示例 2：
//输入: s = "lrloseumgh", k = 6
//输出: "umghlrlose"
//
//限制：
//1 <= k < s.length <= 10000
//
//#

func reverseLeftWords(s string, n int) string {
	length := len(s)
	if length <= n {
		return s
	}

	str := []byte(s)
	str = append(str[n:], str[:n]...)
	return string(str)
}

func reverseLeftWords2(s string, n int) string {
	length := len(s)
	if length <= n {
		return s
	}

	str := []byte(s)
	reverseString(str[:n])
	reverseString(str[n:])
	reverseString(str)

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
