package T000_

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/restore-ip-addresses/

//给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
//
//有效的 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
//
//例如："0.1.2.201" 和 "192.168.1.1" 是 有效的 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效的 IP 地址。
//
//示例 1：
//
//输入：s = "25525511135"
//输出：["255.255.11.135","255.255.111.35"]
//示例 2：

func restoreIpAddresses(input string) (ret []string) {
	var f func(string, []string)

	f = func(s string, arr []string) {
		if len(arr) > 4 {
			return
		}
		if s == "" && len(arr) == 4 {
			ret = append(ret, strings.Join(arr, "."))
			return
		}

		for i := 1; i <= len(s); i++ {
			ip, _ := strconv.Atoi(s[:i])
			if 0 <= ip && ip <= 255 {
				if s[:i][0] == '0' && len(s[:i]) > 1 {
					continue
				}
				arr = append(arr, s[:i])
				f(s[i:], arr)
				arr = arr[:len(arr)-1]
			}
		}
	}

	f(input, nil)
	return
}
