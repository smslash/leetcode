package leetcode

func isAnagram(s string, t string) bool {
	slet := make(map[byte]int)
	tlet := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		slet[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		tlet[t[i]]++
	}

	if len(slet) != len(tlet) {
		return false
	}

	for i, v := range slet {
		if v != tlet[i] {
			return false
		}
	}

	return true
}
