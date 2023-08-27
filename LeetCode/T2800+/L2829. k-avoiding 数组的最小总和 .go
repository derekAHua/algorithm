package LeetCode

func minimumSum(n int, k int) (ret int) {

	m := map[int]struct{}{}
	for i := 1; ; i++ {
		if _, ok := m[k-i]; ok {
			continue
		}
		m[i] = struct{}{}
		if len(m) == n {
			break
		}
	}

	for v := range m {
		ret += v
	}

	return
}
