package T300_

import (
	"container/heap"
	"sort"
)

// https://programmercarl.com/0347.%E5%89%8DK%E4%B8%AA%E9%AB%98%E9%A2%91%E5%85%83%E7%B4%A0.html

//给定一个非空的整数数组，返回其中出现频率前 k 高的元素。
//
//示例 1:
//
//输入: nums = [1,1,1,2,2,3], k = 2
//输出: [1,2]
//示例 2:
//
//输入: nums = [1], k = 1
//输出: [1]
//提示：
//
//你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
//你的算法的时间复杂度必须优于 $O(n \log n)$ , n 是数组的大小。
//题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的。
//你可以按任意顺序返回答案。

// IHeap 构建小顶堆
type (
	IHeap []S
	S     struct {
		value int
		num   int
	}
)

func (h IHeap) Len() int {
	return len(h)
}

func (h IHeap) Less(i, j int) bool {
	return h[i].num < h[j].num
}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(s interface{}) {
	*h = append(*h, s.(S))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) (ret []int) {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	h := new(IHeap)
	heap.Init(h)

	for key, value := range m {
		heap.Push(h, S{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	for i := 0; i < k; i++ {
		ret = append(ret, heap.Pop(h).(S).value)
	}

	return
}

func topKFrequent3(nums []int, k int) (ret []int) {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	for ley := range m {
		ret = append(ret, ley)
	}

	sort.Slice(ret, func(i, j int) bool {
		return m[ret[i]] > m[ret[j]]
	})

	return ret[:k]
}

func topKFrequent2(nums []int, k int) (ret []int) {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	type st struct {
		key, value int
	}
	arr := make([]st, 0, len(m))
	for k, v := range m {
		arr = append(arr, st{k, v})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].value > arr[j].value
	})

	for i := 0; i < k; i++ {
		ret = append(ret, arr[i].key)
	}

	return
}
