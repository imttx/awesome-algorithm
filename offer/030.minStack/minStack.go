package offer

/*
剑指 Offer 30. 包含min函数的栈
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。

示例:
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.min();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.min();   --> 返回 -2.

提示：
各函数的调用总次数不超过 20000 次
*/

/*
本题难点 ： 将 min() 函数复杂度降为 O(1) ，可通过建立辅助栈实现；
数据栈 A ： 栈 A 用于存储所有元素，保证入栈 push() 函数、出栈 pop() 函数、获取栈顶 top() 函数的正常逻辑。
辅助栈 B ： 栈 B 中存储栈 A 中所有 非严格降序 的元素，则栈 A 中的最小元素始终对应栈 B 的栈顶元素，即 min() 函数只需返回栈 B 的栈顶元素即可。
*/

type MinStack struct {
	a []int
	b []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		a: make([]int, 0),
		b: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.a = append(this.a, x)
	if len(this.b) == 0 || this.b[len(this.b)-1] >= x {
		this.b = append(this.b, x)
	}
}

func (this *MinStack) Pop() {
	ak := len(this.a) - 1
	av := this.a[ak]
	this.a = this.a[:ak]
	if av == this.b[len(this.b)-1] {
		this.b = this.b[:len(this.b)-1]
	}
}

func (this *MinStack) Top() int {
	return this.a[len(this.a)-1]
}

func (this *MinStack) Min() int {
	return this.b[len(this.b)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
