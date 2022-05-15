package _0

type MinStack struct {
	min   []int
	stack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		min:   []int{},
		stack: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.min) == 0 {
		this.min = append(this.min, x)
		return
	} else {
		if this.min[len(this.min)-1] < x {
			this.min = append(this.min, this.min[len(this.min)-1])
		} else {
			this.min = append(this.min, x)
		}
		return
	}
}

func (this *MinStack) Pop() {
	this.min = this.min[:len(this.min)-1]
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return -1
	} else {
		return this.stack[len(this.stack)-1]
	}
}

func (this *MinStack) Min() int {
	if len(this.min) == 0 {
		return -1
	} else {
		return this.min[len(this.min)-1]
	}
}
