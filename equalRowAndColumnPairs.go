package leetcode

func equalPairs(grid [][]int) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			equal := true
			for k := 0; k < len(grid); k++ {
				if grid[k][j] != grid[i][k] {
					equal = false
					break
				}
			}

			if equal {
				ans++
			}
		}
	}
	return ans
}
