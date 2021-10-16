package _58_2

func reverseLeftWords(s string, n int) string {
	n = n % len(s)
	if n == 0 {
		return s
	}
	inputBytes := []byte(s)
	outputBytes := make([]byte, len(s))

	for i := 0; i < len(s); i++ {
		outputBytes[(i-n+len(s))%len(s)] = inputBytes[i]
	}
	return string(outputBytes)
}
