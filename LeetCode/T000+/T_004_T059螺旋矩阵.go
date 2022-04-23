package T000_

// https://leetcode-cn.com/problems/spiral-matrix-ii/

/*
给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例:

输入: 3 输出: [ [ 1, 2, 3 ], [ 8, 9, 4 ], [ 7, 6, 5 ] ]
*/

func generateMatrix(n int) (ret [][]int) {

	ret = make([][]int, n)
	for index, _ := range ret {
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
