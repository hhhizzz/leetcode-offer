package _5

func copyArray(target, source []byte) {
	for i := range source {
		target[i] = source[i]
	}
}

func replaceSpace(s string) string {
	inputBytes := []byte(s)
	outputLength := len(inputBytes)
	for _, c := range inputBytes {
		if c == ' ' {
			outputLength += 2
		}
	}
	if outputLength == len(inputBytes) {
		return s
	}
	outputBytes := make([]byte, outputLength)
	outputPos := 0
	for inputPos, c := range inputBytes {
		if c != ' ' {
			outputBytes[outputPos] = inputBytes[inputPos]
			outputPos += 1
		} else {
			copyArray(outputBytes[outputPos:], []byte("%20"))
			outputPos += 3
		}
	}
	return string(outputBytes)
}
