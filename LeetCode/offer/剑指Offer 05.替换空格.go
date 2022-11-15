package offer

// https://leetcode.cn/problems/ti-huan-kong-ge-lcof/

// 请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

func replaceSpace(s string) string {
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
