package T000_

// https://leetcode.cn/problems/ti-huan-kong-ge-lcof/

//请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
//
//示例 1： 输入：s = "We are happy."
//输出："We%20are%20happy."

func replaceSpace(s string) string {
	str := ""
	for _, v := range []byte(s) {
		if v == ' ' {
			str += "%20"
			continue
		}
		str += string(v)
	}
	return str
}

func replaceSpace2(s string) string {
	str := make([]byte, 0, len(s))
	for _, v := range []byte(s) {
		if v == ' ' {
			str = append(str, '%', '2', '0')
			continue
		}
		str = append(str, v)
	}
	return string(str)
}
