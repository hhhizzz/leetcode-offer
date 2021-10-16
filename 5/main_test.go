package _5

import "testing"

func Test_replace(t *testing.T) {
	result := replaceSpace("We are  happy.")
	if result != "We%20are%20%20happy." {
		t.Failed()
	}
}
