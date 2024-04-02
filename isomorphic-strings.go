package leetcode

func isIsomorphic(s string, t string) bool {
	smap := make(map[byte]byte)
	tmap := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		if v1, ok := smap[t[i]]; ok && v1 != s[i] {
			return false
		}
		smap[t[i]] = s[i]

		if v2, ok := tmap[s[i]]; ok && v2 != t[i] {
			return false
		}
		tmap[s[i]] = t[i]
	}

	return true
}
