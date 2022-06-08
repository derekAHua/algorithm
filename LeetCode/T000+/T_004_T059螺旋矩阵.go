package T000_

// https://leetcode-cn.com/problems/spiral-matrix-ii/

func generateMatrix(n int) (ret [][]int) {

	ret = make([][]int, n)
	for index := range ret {
		ret[index] = make([]int, n)
	}

	top, bottom := 0, n-1
	left, right := 0, n-1
	num := 1
	tar := n * n

	for num <= tar {
		for j := left; j <= right; j++ {
			ret[top][j] = num
			num++
		}
		top++
		for i := top; i <= bottom; i++ {
			ret[i][right] = num
			num++
		}
		right--
		for j := right; j >= left; j-- {
			ret[bottom][j] = num
			num++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			ret[i][left] = num
			num++
		}
		left++
	}

	return ret
}

func generateMatrixRepeat1(n int) (ret [][]int) {
	ret = make([][]int, n)
	for i := range ret {
		ret[i] = make([]int, n)
	}

	num, tar := 1, n*n

	top, bottom, left, right := 0, n-1, 0, n-1

	for num < tar {
		for i := left; i < right; i++ {
			ret[top][i] = num
			num++
		}

		for j := top; j < bottom; j++ {
			ret[j][right] = num
			num++
		}

		for i := right; i > left; i-- {
			ret[bottom][i] = num
			num++
		}

		for j := bottom; j > top; j-- {
			ret[j][left] = num
			num++
		}

		top++
		left++
		right--
		bottom--
	}

	if n%2 == 1 {
		ret[n/2][n/2] = num
	}

	return ret
}

func generateMatrixRepeat2(n int) (ret [][]int) {
	ret = make([][]int, n)
	for i := range ret {
		ret[i] = make([]int, n)
	}

	var l, r, t, b = 0, n - 1, 0, n - 1

	max := n * n

	if n%2 == 1 {
		ret[n/2][n/2] = max
	}

	n = 1

	for n < max {
		for i := l; i < r; i++ {
			ret[t][i] = n
			n++
		}
		for j := t; j < b; j++ {
			ret[j][r] = n
			n++
		}
		for i := r; i > l; i-- {
			ret[b][i] = n
			n++
		}
		for j := b; j > t; j-- {
			ret[j][l] = n
			n++
		}

		l++
		r--
		t++
		b--
	}

	return
}
