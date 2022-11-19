package T400_

import "strings"

// https://leetcode.cn/problems/repeated-substring-pattern/

//给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。
//
//示例 1:
//输入: "abab"
//输出: True
//解释: 可由子字符串 "ab" 重复两次构成。
//
//示例 2:
//输入: "aba"
//输出: False
//
//示例 3:
//输入: "abcabcabcabc"
//输出: True
//解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)

func repeatedSubstringPattern2(s string) bool {
	rep := s[1:] + s[:len(s)-1]
	return strings.Contains(rep, s)
}

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}

	next := getNext(s)

	// next[n-1]+1 最长相同前后缀的长度
	if next[n-1] != -1 && n%(n-(next[n-1]+1)) == 0 {
		return true
	}

	return false
}

func getNext(s string) (next []int) {
	next = make([]int, len(s))
	j := -1
	next[0] = j

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
