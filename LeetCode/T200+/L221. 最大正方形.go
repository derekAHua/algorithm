package T200_

// @Author: Derek
// @Description:
// @Date: 2023/9/14 22:57
// @Version 1.0

func maximalSquare(matrix [][]byte) (ret int) {

	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[i]))
	}

	for j := 0; j < len(dp[0]); j++ {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
			ret = 1
		}
	}
	for i := 0; i < len(dp); i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			ret = 1
		}
	}

	m := make(map[int]int)
	all := len(matrix) * len(matrix[0])
	for i := 0; i*i <= all; i++ {
		m[i*i] = i
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if matrix[i][j] == '0' {
				continue
			}
			value := dp[i-1][j-1]
			num := m[value]

			cnt := 0
			for x := 1; x <= num; x++ {
				if matrix[i][j-x] == '1' && matrix[i-x][j] == '1' {
					cnt++
					continue
				} else {
					break
				}
			}

			dp[i][j] = (1 + cnt) * (1 + cnt)
			if ret < dp[i][j] {
				ret = dp[i][j]
			}
		}
	}

	return
}
