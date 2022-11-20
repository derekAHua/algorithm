package T000_

import (
	"testing"
)

// @Author: Derek
// @Description:
// @Date: 2022/11/20 09:57
// @Version 1.0

func Test_combine(t *testing.T) {
	ret := combine(4, 2)
	for _, v := range ret {
		t.Log(v)
	}

	ret = combine(1, 1)
	for _, v := range ret {
		t.Log(v)
	}
}
