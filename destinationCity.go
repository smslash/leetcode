package leetcode

func destCity(paths [][]string) string {
	city := make(map[string]int)
	for i := 0; i < len(paths); i++ {
		for j := 0; j < 2; j++ {
			if j == 0 {
				city[paths[i][j]] += 2
			} else {
				city[paths[i][j]]++
			}
		}
	}

	for i, v := range city {
		if v == 1 {
			return i
		}
	}

	return ""
}
