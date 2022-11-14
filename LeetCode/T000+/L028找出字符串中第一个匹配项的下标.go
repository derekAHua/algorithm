package T000_

// https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/

//实现 strStr() 函数。
//
//给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
//
//示例 1: 输入: haystack = "hello", needle = "ll" 输出: 2
//
//示例 2: 输入: haystack = "aaaaa", needle = "bba" 输出: -1
//
//说明: 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。 对于本题而言，当 needle 是空字符串时我们应当返回 0 。
//这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。

func strStr1(haystack string, needle string) int {
	needLen := len(needle)
	length := len(haystack)
	if needLen > length {
		return -1
	}
	if len(needle) == 0 {
		return 0
	}

	for index := 0; index < length-needLen+1; index++ {
		if haystack[index:index+needLen] == needle {
			return index
		}
	}

	return -1
}

func strStr2(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	next := getNext(needle)
	j := -1 // 模式串的起始位置 next为-1 因此也为-1
	for i := 0; i < len(haystack); i++ {
		for j >= 0 && haystack[i] != needle[j+1] {
			j = next[j] // 寻找下一个匹配点
		}
		if haystack[i] == needle[j+1] {
			j++
		}
		if j == len(needle)-1 { // j指向了模式串的末尾
			return i - len(needle) + 1
		}
	}
	return -1
}

func getNext(s string) (next []int) {
	next = make([]int, len(s))
	j := -1
	next[0] = -1
	// i:前缀的末尾位置
	// j:后缀的末尾位置;最长相等前后缀长度

	for i := 1; i < len(s); i++ {
		for j >= 0 && s[i] != s[j+1] {
			j = next[j]
		}
		if s[i] == s[j+1] {
			j++
		}
		next[i] = j
	}

	return
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	next := get(needle)
	j := -1
	for i := 0; i < len(haystack); i++ {
		for j >= 0 && haystack[i] != needle[j+1] {
			j = next[j]
		}
		if haystack[i] == needle[j+1] {
			j++
		}
		if j == len(needle)-1 {
			return i - len(needle) + 1
		}
	}
	return -1
}

func get(s string) []int {
	ret := make([]int, len(s))
	j := -1
	ret[0] = j

	for i := 1; i < len(s); i++ {
		if j >= 0 && s[i] != s[j+1] {
			j = ret[j]
		}
		if s[i] == s[j+1] {
			j++
		}
		ret[i] = j
	}
	return ret
}
