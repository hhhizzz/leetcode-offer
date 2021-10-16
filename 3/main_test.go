package _3

import "testing"

func TestFindRepeatNumber(t *testing.T) {
	var result int

	result = findRepeatNumber([]int{3, 4, 2, 0, 0, 1})
	if result != 0 {
		t.Fail()
	}
}
