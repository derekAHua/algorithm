package T000_

import (
	"testing"
)

// @Author: Derek
// @Description:
// @Date: 2022/11/10 21:59
// @Version 1.0

func Test_swapPairs(t *testing.T) {
	h := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}

	swapPairs(h)

}
