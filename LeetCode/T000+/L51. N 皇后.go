package T000_

import "strings"

// https://leetcode.cn/problems/n-queens/

func solveNQueens(n int) (ret [][]string) {
	chessboard := make([][]string, n)
	for i := 0; i < n; i++ {
		chessboard[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			chessboard[i][j] = "."
		}
	}

	var f func(int)
	f = func(row int) {
		if row == n {
			temp := make([]string, n)
			for i, rowStr := range chessboard {
				temp[i] = strings.Join(rowStr, "")
			}
			ret = append(ret, temp)
			return
		}

		for i := 0; i < n; i++ {
			if can(n, row, i, chessboard) {
				chessboard[row][i] = "Q"
				f(row + 1)
				chessboard[row][i] = "."
			}
		}
	}

	f(0)
	return
}

func can(n, row, col int, chessboard [][]string) bool {
	for i := 0; i < row; i++ {
		if chessboard[i][col] == "Q" {
			return false
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}

	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}
	return true
}
