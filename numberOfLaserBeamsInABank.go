package leetcode

func numberOfBeams(bank []string) int {
	res, count, prev := 0, 0, 0
	for i := len(bank) - 1; i > -1; i-- {
		for j := 0; j < len(bank[i]); j++ {
			if bank[i][j] == '1' {
				res += prev
			}
		}
		found := false
		for j := 0; j < len(bank[i]); j++ {
			if bank[i][j] == '1' {
				count++
				found = true
			}
		}
		if found {
			prev = count
			count = 0
		}
	}
	return res
}
