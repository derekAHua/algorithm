package T000_

import (
	"testing"
)

// @Author: Derek
// @Description:
// @Date: 2022/11/20 15:15
// @Version 1.0

func Test_subsets(t *testing.T) {
	ret := subsets([]int{1, 2, 3})
	for _, v := range ret {
		t.Log(v)
	}
}
