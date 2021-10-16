package _58_2

import "testing"

func TestReverseLeftWords(t *testing.T) {
	var result string

	result = reverseLeftWords("abcdefg", 2)
	if result != "cdefgab" {
		t.Failed()
	}

	result = reverseLeftWords("lrloseumgh", 6)
	if result != "umghlrlose" {
		t.Failed()
	}
}
