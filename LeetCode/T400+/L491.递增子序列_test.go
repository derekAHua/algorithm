package T400_

import (
	"testing"
)

// @Author: Derek
// @Description:
// @Date: 2022/11/20 16:36
// @Version 1.0

func Test_findSubsequences(t *testing.T) {
	subsequences := findSubsequences([]int{4, 4, 3, 2, 1})
	for _, v := range subsequences {
		t.Log(v)
	}
}
