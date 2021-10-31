package _11

import "testing"

func TestMin(t *testing.T) {
	var input []int
	var output int

	input = []int{1, 2, 3, 4, 5}
	output = minArray(input)
	if output != 1 {
		t.Error(output)
	}

	input = []int{2, 3, 4, 5, 1}
	output = minArray(input)
	if output != 1 {
		t.Error(output)
	}

	input = []int{5, 1, 2, 3, 4}
	output = minArray(input)
	if output != 1 {
		t.Error(output)
	}

	input = []int{3,1,3}
	output = minArray(input)
	if output != 1 {
		t.Error(output)
	}
}
