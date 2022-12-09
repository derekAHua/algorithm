package T800_

// https://leetcode.cn/problems/lemonade-change/

func lemonadeChange(bills []int) (ret bool) {
	if len(bills) == 0 {
		return true
	}
	if bills[0] != 5 {
		return
	}

	m := map[int]int{
		5: 1,
	}
	for i := 1; i < len(bills); i++ {
		if m[5] < 0 {
			return false
		}

		switch bills[i] {
		case 5:
			m[5]++
		case 10:
			m[10]++
			m[5]--
		case 20:
			if m[10] > 0 {
				m[10]--
				m[5]--
				continue
			}
			m[5] -= 3
		}
	}

	if m[5] < 0 {
		return false
	}

	return true
}
