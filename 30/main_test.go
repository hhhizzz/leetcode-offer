package _30

import "testing"

func TestStack(t *testing.T) {
	minStack := Constructor()

	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	if minStack.Min() != -3 {
		t.Fatal()
	} //  --> 返回 -3.
	minStack.Pop()
	if minStack.Top() != 0 {
		t.Fatal()
	} // --> 返回 0.
	if minStack.Min() != -2 {
		t.Fatal()
	} //--> 返回 -2.

}
