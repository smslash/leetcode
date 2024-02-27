package leetcode

func dayOfYear(date string) int {
	year, month, day := splitDate(date)
	leap := isLeapYear(year)
	res := day

	for i := 1; i < month; i++ {
		if i == 1 {
			res += 31
		} else if i == 2 {
			if leap {
				res += 29
			} else {
				res += 28
			}
		} else if i == 3 {
			res += 31
		} else if i == 4 {
			res += 30
		} else if i == 5 {
			res += 31
		} else if i == 6 {
			res += 30
		} else if i == 7 {
			res += 31
		} else if i == 8 {
			res += 31
		} else if i == 9 {
			res += 30
		} else if i == 10 {
			res += 31
		} else if i == 11 {
			res += 30
		} else if i == 12 {
			res += 31
		}

	}

	return res
}

func isLeapYear(n int) bool {
	if n%400 == 0 || (n%4 == 0 && n%100 != 0) {
		return true
	}
	return false
}

func splitDate(s string) (int, int, int) {
	y, m, d := "", "", ""
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '-' {
			count++
		} else if count == 0 {
			y += string(s[i])
		} else if count == 1 {
			m += string(s[i])
		} else if count == 2 {
			d += string(s[i])
		}
	}

	ny, nm, nd := 0, 0, 0

	ny, _ = strconv.Atoi(y)
	nm, _ = strconv.Atoi(m)
	nd, _ = strconv.Atoi(d)

	return ny, nm, nd
}
