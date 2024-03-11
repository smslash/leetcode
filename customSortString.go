package leetcode

func customSortString(order string, s string) string {
	res := []byte{}
	m := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	for i := 0; i < len(order); i++ {
		if _, ok := m[order[i]]; ok {
			for m[order[i]] > 0 {
				res = append(res, order[i])
				m[order[i]]--
			}
		}
	}

	for i, v := range m {
		for v > 0 {
			res = append(res, i)
			v--
		}
	}

	return string(res)
}
