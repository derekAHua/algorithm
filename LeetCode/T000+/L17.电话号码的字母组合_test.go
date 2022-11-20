package T000_

import (
	"testing"
)

// @Author: Derek
// @Description:
// @Date: 2022/11/20 10:51
// @Version 1.0

func Test_letterCombinations(t *testing.T) {
	gotRet := letterCombinations("23")
	for _, v := range gotRet {
		t.Log(v)
	}
}
