package T300_

import "testing"

// @Author: Derek
// @Description:
// @Date: 2022/11/12 23:04
// @Version 1.0

func Test_reverseString(t *testing.T) {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	t.Log(string(s) == "olleh")
}
