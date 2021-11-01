package _50

func firstUniqChar(s string) byte {
	m := make([]int, 26)
	var q []int

	for _, c := range s{
		pos := int(c - 'a')
		m[pos] += 1
		if m[pos] == 1{
			q = append(q, pos)
		}
		for len(q) > 0 && m[q[0]] > 1{
			q = q[1:]
		}
	}
	if len(q) != 0{
		return byte(q[0]+'a')
	}
	return ' '
}
