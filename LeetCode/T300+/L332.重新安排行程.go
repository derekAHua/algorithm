package T300_

import (
	"sort"
)

// https://leetcode.cn/problems/reconstruct-itinerary/

//给定一个机票的字符串二维数组 [from, to]，子数组中的两个成员分别表示飞机出发和降落的机场地点，对该行程进行重新规划排序。
//所有这些机票都属于一个从 JFK（肯尼迪国际机场）出发的先生，所以该行程必须从 JFK 开始。
//
//提示：
//
//如果存在多种有效的行程，请你按字符自然排序返回最小的行程组合。例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"] 相比就更小，排序更靠前
//所有的机场都用三个大写字母表示（机场代码）。
//假定所有机票至少存在一种合理的行程。
//所有的机票必须都用一次 且 只能用一次。
//示例 1：
//
//输入：[["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]]
//输出：["JFK", "MUC", "LHR", "SFO", "SJC"]
//示例 2：
//
//输入：[["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
//输出：["JFK","ATL","JFK","SFO","ATL","SFO"]
//解释：另一种有效的行程是 ["JFK","SFO","ATL","JFK","ATL","SFO"]。但是它自然排序更大更靠后。

type pair struct {
	target  string
	visited bool
}
type pairs []*pair

func findItinerary(tickets [][]string) (ret []string) {
	// map[出发机场] pair{目的地,是否被访问过}
	targets := make(map[string]pairs)
	for _, ticket := range tickets {
		targets[ticket[0]] = append(targets[ticket[0]], &pair{target: ticket[1], visited: false})
	}
	for k := range targets {
		sort.Slice(targets[k], func(i, j int) bool {
			return targets[k][i].target < targets[k][j].target
		})
	}
	ret = append(ret, "JFK")

	var backtracking func() bool

	backtracking = func() bool {
		if len(tickets)+1 == len(ret) {
			return true
		}
		// 取出起飞航班对应的目的地
		for _, pair := range targets[ret[len(ret)-1]] {
			if pair.visited == false {
				ret = append(ret, pair.target)
				pair.visited = true
				if backtracking() {
					return true
				}
				ret = ret[:len(ret)-1]
				pair.visited = false
			}
		}
		return false
	}

	backtracking()

	return
}
