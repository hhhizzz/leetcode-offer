package _50

func firstUniqChar(s string) byte {
	m := make([]int, 26)
	var order []int

	for _,c32 := range s{
		pos := int(c32 - 'a')
		m[pos] += 1
		if m[pos] == 1{
			order = append(order,pos)
		}
	}

	for _,c := range order{
		if m[c] == 1{
			return byte(c+'a')
		}
	}
	return ' '
}
