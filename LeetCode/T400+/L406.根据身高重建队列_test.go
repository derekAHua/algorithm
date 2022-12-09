package T400_

import (
	"testing"
)

// @Author: Derek
// @Description:
// @Date: 2022/12/3 15:36
// @Version 1.0

func Test_reconstructQueue(t *testing.T) {
	queue := reconstructQueue([][]int{
		{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2},
	})

	t.Log(queue)
}
