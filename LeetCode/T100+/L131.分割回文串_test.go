package T100_

import (
	"testing"
)

// @Author: Derek
// @Description:
// @Date: 2022/11/20 14:35
// @Version 1.0

func Test_partition(t *testing.T) {
	gotRet := partition("aab")
	for _, v := range gotRet {
		t.Log(v)
	}
}
