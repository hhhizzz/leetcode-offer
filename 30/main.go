package _30

type MinStack struct {
	data    []int
	minData []int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data:    []int{},
		minData: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.minData) == 0 {
		this.minData = append(this.minData, x)
	} else {
		currentMin := this.minData[len(this.minData)-1]
		if x <= currentMin {
			this.minData = append(this.minData, x)
		}
	}
}

func (this *MinStack) Pop() {
	if this.Top() == this.Min() {
		minLength := len(this.minData)
		this.minData = this.minData[:minLength-1]
	}
	length := len(this.data)
	this.data = this.data[:length-1]
}

func (this *MinStack) Top() int {
	length := len(this.data)
	return this.data[length-1]
}

func (this *MinStack) Min() int {
	length := len(this.minData)
	return this.minData[length-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
