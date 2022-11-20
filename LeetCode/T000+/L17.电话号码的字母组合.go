package T000_

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/

//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
//
//给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//示例: 输入："23" 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
//
//说明：尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。

func letterCombinations(digits string) (ret []string) {
	length := len(digits)
	if length == 0 {
		return
	}

	m := [10]string{
		"",     // 0
		"",     // 1
		"abc",  // 2
		"def",  // 3
		"ghi",  // 4
		"jkl",  // 5
		"mno",  // 6
		"pqrs", // 7
		"tuv",  // 8
		"wxyz", // 9
	}

	var f func(int, []byte)

	f = func(start int, str []byte) {
		if start >= length {
			ret = append(ret, string(str))
			return
		}

		cur := m[digits[start]-'0']
		for _, v := range cur {
			str = append(str, byte(v))
			f(start+1, str)
			str = str[:len(str)-1]
		}
	}

	f(0, nil)

	return
}
