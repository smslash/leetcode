package leetcode

func findMissingAndRepeatedValues(grid [][]int) []int {
	var (
		n    = len(grid)
		m    = make(map[int]int)
		a, b = 0, 0
	)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			m[grid[i][j]]++
			if m[grid[i][j]] >= 2 {
				a = grid[i][j]
			}
		}
	}

	for i := 1; i <= n*n; i++ {
		if m[i] == 0 {
			return []int{a, i}
		}
	}

	return []int{a, b}
}
