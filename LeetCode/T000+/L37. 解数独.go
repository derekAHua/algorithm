package T000_

// https://leetcode.cn/problems/sudoku-solver/

func solveSudoku(board [][]byte) {
	var backtracking func(board [][]byte) bool

	backtracking = func(board [][]byte) bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				// 判断此位置是否适合填数字
				if board[i][j] != '.' {
					continue
				}
				// 尝试填1-9
				for k := '1'; k <= '9'; k++ {
					if canValid(i, j, byte(k), board) == true { // 如果满足要求就填
						board[i][j] = byte(k)
						if backtracking(board) == true {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
		return true
	}

	backtracking(board)
}

// 判断填入数字是否满足要求
func canValid(row, col int, k byte, board [][]byte) bool {
	for i := 0; i < 9; i++ { // 行
		if board[row][i] == k {
			return false
		}
	}

	for i := 0; i < 9; i++ { // 列
		if board[i][col] == k {
			return false
		}
	}

	// 方格
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == k {
				return false
			}
		}
	}
	return true
}
