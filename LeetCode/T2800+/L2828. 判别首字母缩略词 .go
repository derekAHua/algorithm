package LeetCode

func isAcronym(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}

	idx := 0
	for _, v := range words {
		if len(v) == 0 {
			return false
		}
		if idx == len(s) {
			return false
		}
		if v[0] != s[idx] {
			return false
		}
		idx++
	}

	return true
}
