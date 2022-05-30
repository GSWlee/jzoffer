package _9

type MaxQueue struct {
	dq  []int
	max []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		dq:  []int{},
		max: []int{},
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.dq) == 0 {
		return -1
	}
	return this.max[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.dq = append(this.dq, value)
	for len(this.max) > 0 {
		if this.max[len(this.max)-1] < value {
			this.max = this.max[:len(this.max)-1]
		} else {
			break
		}
	}
	this.max = append(this.max, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.dq) == 0 {
		return -1
	}
	temp := this.dq[0]
	this.dq = this.dq[1:]
	if temp == this.max[0] {
		this.max = this.max[1:]
	}
	return temp
}
