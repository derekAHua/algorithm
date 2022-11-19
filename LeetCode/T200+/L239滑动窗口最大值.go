package T200_

// https://programmercarl.com/0239.%E6%BB%91%E5%8A%A8%E7%AA%97%E5%8F%A3%E6%9C%80%E5%A4%A7%E5%80%BC.html

//给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
//你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
//
//返回滑动窗口中的最大值。

type Queue struct {
	s []int
}

func (q *Queue) Front() int {
	return q.s[0]
}

func (q *Queue) Push(val int) {
	for !q.Empty() && val > q.Back() {
		q.s = q.s[:len(q.s)-1]
	}
	q.s = append(q.s, val)
}

func (q *Queue) Empty() bool {
	return len(q.s) == 0
}

func (q *Queue) Pop(val int) {
	if q.Empty() {
		return
	}

	if q.s[0] == val {
		q.s = q.s[1:]
	}

	return
}

func (q *Queue) Back() int {
	return q.s[len(q.s)-1]
}

func maxSlidingWindow(nums []int, k int) (ret []int) {
	q := Queue{}
	for i := 0; i < len(nums) && i < k; i++ {
		q.Push(nums[i])
	}

	ret = append(ret, q.Front())

	for i := k; i < len(nums); i++ {
		left := i - k
		q.Pop(nums[left])
		q.Push(nums[i])
		ret = append(ret, q.Front())
	}

	return
}
