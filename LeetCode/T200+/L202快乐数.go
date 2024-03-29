package T200_

// https://leetcode.cn/problems/happy-number/

//编写一个算法来判断一个数 n 是不是快乐数。
//
//「快乐数」定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。如果 可以变为  1，那么这个数就是快乐数。
//
//如果 n 是快乐数就返回 True ；不是，则返回 False 。

func isHappy(n int) bool {
	m := make(map[int]bool)
	for n != 1 && !m[n] {
		m[n] = true
		n = getSum(n)
	}

	return n == 1
}

func getSum(n int) (ret int) {
	for n != 0 {
		ret += (n % 10) * (n % 10)
		n /= 10
	}
	return
}
