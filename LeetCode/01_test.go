package LeetCode

import "testing"

// @Author: Derek
// @Description:
// @Date: 2022/12/10 11:32
// @Version 1.0

func Test_weiBagProblem(t *testing.T) {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	println(weiBagProblem(weight, value, 4))
	println(weiBagProblem2(weight, value, 4))
}
