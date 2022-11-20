package T000_

import (
	"testing"
)

// @Author: Derek
// @Description:
// @Date: 2022/11/20 18:41
// @Version 1.0

func Test_permuteUnique(t *testing.T) {
	unique := permuteUnique([]int{1, 1, 2})
	for _, v := range unique {
		t.Log(v)
	}
}
