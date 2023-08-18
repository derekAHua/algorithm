package T100_

// https://leetcode.cn/problems/symmetric-tree/

//给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
//
//返回 s 所有可能的分割方案。
//
//示例: 输入: "aab" 输出: [ ["aa","b"], ["a","a","b"] ]

func partition(s string) (ret [][]string) {
	var arr []string
	var f func(idx int)
	f = func(idx int) {
		if idx == len(s) {
			ret = append(ret, append([]string{}, arr...))
			return
		}

		for i := idx + 1; i <= len(s); i++ {
			vs := s[idx:i]
			if !is(vs) {
				continue
			}
			arr = append(arr, vs)
			f(i)
			arr = arr[:len(arr)-1]
		}
	}

	f(0)
	return
}

func partitionR1(s string) (ret [][]string) {
	var f func(string, []string)

	f = func(s string, cur []string) {
		if s == "" {
			r := make([]string, len(cur))
			copy(r, cur)
			ret = append(ret, r)
			return
		}

		for i := 1; i <= len(s); i++ {
			if is(s[:i]) {
				cur = append(cur, s[:i])
				f(s[i:], cur)
				cur = cur[:len(cur)-1]
			}
		}
	}

	f(s, nil)

	return
}

func is(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
