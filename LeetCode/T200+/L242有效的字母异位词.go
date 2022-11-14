package T200_

// https://leetcode.cn/problems/valid-anagram/

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
//示例 1: 输入: s = "anagram", t = "nagaram" 输出: true
//
//示例 2: 输入: s = "rat", t = "car" 输出: false
//
//说明: 你可以假设字符串只包含小写字母。

func isAnagram(s string, t string) bool {
	var arr [26]int
	for _, v := range s {
		arr[v-'a']++
	}
	for _, v := range t {
		arr[v-'a']--
	}

	return arr == [26]int{}
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m1 := make(map[int32]int)
	m2 := make(map[int32]int)

	for _, k := range s {
		m1[k] = m1[k] + 1
	}

	for _, k := range t {
		m2[k] = m2[k] + 1
	}

	for k := range m1 {
		if m1[k] != m2[k] {
			return false
		}
	}

	return true
}
