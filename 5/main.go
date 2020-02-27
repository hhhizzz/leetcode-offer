package _5

func replaceSpace(s string) string {
    bytes := []byte(s)
    size := len(bytes)
    for i := 0; i < len(bytes); i++ {
        if bytes[i] == ' ' {
            size += 2
        }
    }
    newBytes := make([]byte, size)
    for i, j := 0, 0; i < len(bytes); i++ {
        if bytes[i] != ' ' {
            newBytes[j] = bytes[i]
            j++
        } else {
            newBytes[j] = '%'
            newBytes[j+1] = '2'
            newBytes[j+2] = '0'
            j += 3
        }
    }
    return string(newBytes)
}
