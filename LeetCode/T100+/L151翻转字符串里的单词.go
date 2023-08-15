package T100_

import "strings"

// https://leetcode.cn/problems/reverse-words-in-a-string/

//给定一个字符串，逐个翻转字符串中的每个单词。
//
//示例 1：
//输入: "the sky is blue"
//输出: "blue is sky the"
//
//示例 2：
//输入: "  hello world!  "
//输出: "world! hello"
//解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
//
//示例 3：
//输入: "a good   example"
//输出: "example good a"
//解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

func reverseWords(s string) (ret string) {
	str := make([]byte, 0, len(s))
	for i, v := range s {
		if v == ' ' {
			if i == 0 || s[i-1] == ' ' {
				continue
			}
		}
		str = append(str, byte(v))
	}

	for i := range str {
		if str[i] != ' ' {
			str = str[i:]
			break
		}
	}

	for i := len(str) - 1; i >= 0; i-- {
		if str[i] != ' ' {
			str = str[:i+1]
			break
		}
	}

	split := strings.Split(string(str), " ")
	for i := len(split) - 1; i >= 0; i-- {
		ret += split[i] + " "
	}
	return ret[:len(ret)-1]
}

func reverseWordsR1(s string) string {
	ret := ""
	split := strings.Split(s, " ")
	for i := len(split) - 1; i >= 0; i-- {
		if split[i] == "" {
			continue
		}
		ret += split[i] + " "
	}
	return ret[:len(ret)-1]
}

func reverseWords2(s string) string {
	if len(s) == 0 {
		return s
	}

	//1.使用双指针删除冗余的空格
	slowIndex, fastIndex := 0, 0
	b := []byte(s)
	//删除头部冗余空格
	for fastIndex < len(b) && b[fastIndex] == ' ' {
		fastIndex++
	}
	//删除单词间冗余空格
	for ; fastIndex < len(b); fastIndex++ {
		if fastIndex-1 > 0 && b[fastIndex-1] == b[fastIndex] && b[fastIndex] == ' ' {
			continue
		}
		b[slowIndex] = b[fastIndex]
		slowIndex++
	}
	//删除尾部冗余空格
	if slowIndex-1 > 0 && b[slowIndex-1] == ' ' {
		b = b[:slowIndex-1]
	} else {
		b = b[:slowIndex]
	}

	reverseString(b)
	left := 0
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' {
			reverseString(b[left:i])
			left = i + 1
		}
	}
	reverseString(b[left:])

	return string(b)
}

func reverseWords3(s string) string {
	length := len(s)
	if length == 0 {
		return s
	}
	bytes := []byte(s)

	slow, fast := 0, 0
	for fast < length && bytes[fast] == ' ' {
		fast++
	}

	for ; fast < length; fast++ {
		if fast-1 >= 0 && bytes[fast] == ' ' && bytes[fast-1] == ' ' {
			continue
		}
		bytes[slow] = bytes[fast]
		slow++
	}

	if slow-1 > 0 && bytes[slow-1] == ' ' {
		bytes = bytes[:slow-1]
	} else {
		bytes = bytes[:slow]
	}

	reverseString(bytes)

	cur := 0
	for i, v := range bytes {
		if v == ' ' {
			reverseString(bytes[cur:i])
			cur = i + 1
		}
	}
	reverseString(bytes[cur:])

	return string(bytes)
}

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
