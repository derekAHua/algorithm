package T500_

// https://leetcode.cn/problems/fibonacci-number/

func fib(n int) int {
	if n <= 1 {
		return n
	}

	arr := make([]int, n+1)
	arr[0] = 0
	arr[1] = 1
	for i := 2; i < n+1; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr[n]
}

func fib2(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
