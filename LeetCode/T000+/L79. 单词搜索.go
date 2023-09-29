package T000_

// @Author: Derek
// @Description:
// @Date: 2023/9/27 22:00
// @Version 1.0

func exist(arr [][]byte, target string) (ret bool) {
	done := make([][]bool, len(arr))
	for i := range done {
		done[i] = make([]bool, len(arr[i]))
	}

	var f func(i, j, idx int) bool
	f = func(i, j, idx int) bool {
		if idx == len(target) {
			return true
		}
		if i < 0 || j < 0 || i == len(arr) || j == len(arr[i]) {
			return false
		}
		if done[i][j] || arr[i][j] != target[idx] {
			return false
		}

		done[i][j] = true
		defer func() { done[i][j] = false }()
		r1 := f(i-1, j, idx+1)
		r2 := f(i+1, j, idx+1)
		r3 := f(i, j-1, idx+1)
		r4 := f(i, j+1, idx+1)
		return r1 || r2 || r3 || r4
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if f(i, j, 0) {
				return true
			}
		}
	}

	return
}
