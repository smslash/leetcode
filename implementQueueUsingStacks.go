package leetcode

type MyQueue struct {
	q []int
}

func Constructor() MyQueue {
	return MyQueue{[]int{}}
}

func (this *MyQueue) Push(x int) {
	this.q = append(this.q, x)
}

func (this *MyQueue) Pop() int {
	res := this.q[0]
	this.q = this.q[1:]
	return res
}

func (this *MyQueue) Peek() int {
	return this.q[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.q) == 0
}
