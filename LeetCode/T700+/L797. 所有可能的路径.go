package T700_

// @Author: Derek
// @Description:
// @Date: 2023/9/3 15:00
// @Version 1.0

func allPathsSourceTarget(graph [][]int) (ret [][]int) {

	var arr []int
	var f func(idx int)
	f = func(idx int) {
		if idx == len(graph)-1 {
			ret = append(ret, append([]int{}, arr...))
			return
		}

		for _, v := range graph[idx] {
			arr = append(arr, v)
			f(v)
			arr = arr[:len(arr)-1]
		}
	}

	arr = append(arr, 0)
	f(0)
	return
}
