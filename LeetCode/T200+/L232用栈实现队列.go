package T200_

// https://leetcode.cn/problems/implement-queue-using-stacks/

//使用栈实现队列的下列操作：
//
//push(x) -- 将一个元素放入队列的尾部。
//pop() -- 从队列首部移除元素。
//peek() -- 返回队列首部的元素。
//empty() -- 返回队列是否为空。

type MyQueue struct {
	in  []int
	out []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyQueue) Pop() (ret int) {
	l := len(this.out)
	if l == 0 {
		this.out = this.in
		this.in = nil
	}

	ret = this.out[0]
	this.out = this.out[1:]
	return
}

func (this *MyQueue) Peek() int {
	if len(this.out) == 0 && len(this.in) > 0 {
		return this.in[0]
	}
	return this.out[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.out) == 0 && len(this.in) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
