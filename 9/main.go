package _9

type CStack struct {
	data []int
}

func (this *CStack) AppendTail(value int) {
	this.data = append(this.data, value)
}

func (this *CStack) DeleteTail() int {
	currentLen := len(this.data)
	var tail = this.data[currentLen-1]
	this.data = this.data[:currentLen-1]
	return tail
}

func (this *CStack) Len() int {
	return len(this.data)
}

type CQueue struct {
	activeStack CStack
	helpStack   CStack
}

func MoveToAnotherStack(s1 *CStack, s2 *CStack) {
	value := s1.DeleteTail()
	s2.AppendTail(value)
}

func Constructor() CQueue {
	return CQueue{
		activeStack: CStack{[]int{}},
		helpStack:   CStack{[]int{}},
	}
}

func (this *CQueue) AppendTail(value int) {
	this.activeStack.AppendTail(value)
}

func (this *CQueue) DeleteHead() int {
	if this.helpStack.Len() == 0 {
		for this.activeStack.Len() > 0 {
			MoveToAnotherStack(&this.activeStack, &this.helpStack)
		}
	}
	if this.helpStack.Len() == 0 {
		return -1
	}
	return this.helpStack.DeleteTail()
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
