package T200_

import (
	"testing"
)

// @Author: Derek
// @Description:
// @Date: 2022/11/20 10:17
// @Version 1.0

func Test_combinationSum3(t *testing.T) {
	ret := combinationSum3(3, 7)
	for _, v := range ret {
		t.Log(v)
	}

	ret = combinationSum3(9, 45)
	for _, v := range ret {
		t.Log(v)
	}
}
