package leetcode

func countCharacters(words []string, chars string) int {
	f := make([]int, 26)
	for _, r := range chars {
		f[r-'a']++
	}

	ff := func() []int {
		clF := make([]int, 26)
		for i, v := range f {
			clF[i] = v
		}
		return clF
	}

	lt := 0
	ok := true
	for _, w := range words {
		lw := len(w)
		ok = true
		cff := ff()
		for _, rw := range w {
			if cff[rw-'a'] == 0 {
				ok = false
				break
			} else {
				cff[rw-'a']--
			}
		}
		if ok {
			lt += lw
		}
	}

	return lt
}
