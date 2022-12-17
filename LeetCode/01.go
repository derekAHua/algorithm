package LeetCode

func weiBagProblem(weight, value []int, bagWeight int) int {
	dp := make([]int, bagWeight+1)
	for index, we := range weight {
		for i := bagWeight; i >= we; i-- {
			dp[i] = max(dp[i], dp[i-we]+value[index])
		}

	}

	return dp[bagWeight]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func weiBagProblem2(weight, value []int, bagWeight int) int {
	// 定义 and 初始化
	dp := make([]int, bagWeight+1)
	// 递推顺序
	for i := 0; i < len(weight); i++ {
		// 这里必须倒序,区别二维,因为二维dp保存了i的状态
		for j := bagWeight; j >= weight[i]; j-- {
			// 递推公式
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	//fmt.Println(dp)
	return dp[bagWeight]
}
