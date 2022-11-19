package T000_

// https://leetcode.cn/problems/valid-parentheses/

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//注意空字符串可被认为是有效字符串。
//示例 1:
//
//输入: "()"
//输出: true
//示例 2:
//
//输入: "()[]{}"
//输出: true
//示例 3:
//
//输入: "(]"
//输出: false
//示例 4:
//
//输入: "([)]"
//输出: false

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	stack := make([]int32, 0)
	for _, v := range s {
		if len(stack) == 0 {
			stack = append(stack, v)
			continue
		}

		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
			continue
		}

		top := stack[len(stack)-1]
		if !check(top, v) {
			return false
		}
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}

func check(u int32, value int32) bool {
	return (u == '(' && value == ')') ||
		(u == '[' && value == ']') ||
		(u == '{' && value == '}')
}
