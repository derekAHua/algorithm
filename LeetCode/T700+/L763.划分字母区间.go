package T700_

// https://leetcode.cn/problems/partition-labels/

func partitionLabels(s string) (ret []int) {
	if len(s) == 0 {
		return
	}

	var arr [26]int
	for i, v := range s {
		arr[v-'a'] = i
	}

	var left, right int
	for i := 0; i < len(s); i++ {
		if arr[s[i]-'a'] > right {
			right = arr[s[i]-'a']
		}

		if i == right {
			ret = append(ret, right-left+1)
			left = i + 1
		}
	}

	return
}
