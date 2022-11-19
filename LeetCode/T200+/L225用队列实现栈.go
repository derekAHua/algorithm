package T200_

// https://leetcode.cn/problems/implement-stack-using-queues/

//使用队列实现栈的下列操作：
//
//push(x) -- 元素 x 入栈
//pop() -- 移除栈顶元素
//top() -- 获取栈顶元素
//empty() -- 返回栈是否为空
//注意:
//
//你只能使用队列的基本操作-- 也就是 push to back, peek/pop from front, size, 和 is empty 这些操作是合法的。
//你所使用的语言也许不支持队列。 你可以使用 list 或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。
//你可以假设所有操作都是有效的（例如, 对一个空的栈不会调用 pop 或者 top 操作）。

type MyStack struct {
	in []int
}

func Constructor2() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.in = append([]int{x}, this.in...)
}

func (this *MyStack) Pop() (ret int) {
	if len(this.in) == 0 {
		return 0
	}

	ret = this.in[0]
	this.in = this.in[1:]
	return ret
}

func (this *MyStack) Top() int {
	if len(this.in) == 0 {
		return 0
	}

	return this.in[0]
}

func (this *MyStack) Empty() bool {
	return len(this.in) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
