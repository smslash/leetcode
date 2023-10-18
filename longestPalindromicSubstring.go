package leetcode

func evenPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s)-1; i++ {
		l, r := i, i+1

		if s[l] != s[r] {
			continue
		}

		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > len(res) {
				res = s[l : r+1]
			}
			l--
			r++
		}

	}
	return res
}

func oddPalindrome(s string) string {
	res := ""
	if len(s) > 0 {
		res = string(s[0])
	}
	for i := 0; i < len(s)-2; i++ {
		l, r := i, i+2

		if s[l] != s[r] {
			continue
		}

		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > len(res) {
				res = s[l : r+1]
			}
			l--
			r++
		}

	}
	return res
}

func longestPalindrome(s string) string {
	evenRes := evenPalindrome(s)
	oddRes := oddPalindrome(s)

	if len(evenRes) > len(oddRes) {
		return evenRes
	} else {
		return oddRes
	}
}
