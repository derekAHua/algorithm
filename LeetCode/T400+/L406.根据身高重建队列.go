package T400_

import "sort"

// https://leetcode.cn/problems/queue-reconstruction-by-height/

func reconstructQueue(people [][]int) [][]int {
	// 先将身高从大到小排序，确定最大个子的相对位置
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1] // 这个才是当身高相同时，将K按照从小到大排序
		}
		return people[i][0] > people[j][0] // 这个只是确保身高按照由大到小的顺序来排，并不确定K是按照从小到大排序的
	})

	// 再按照K进行插入排序，优先插入K小的
	result := make([][]int, 0)
	for _, info := range people {
		result = append(result, info)

		k := info[1]
		copy(result[k+1:], result[k:]) // 将插入位置之后的元素后移动一位（意思是腾出空间）
		result[k] = info               // 将插入元素位置插入元素
	}

	return result
}
