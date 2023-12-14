package leetcode

func onesMinusZeros(grid [][]int) [][]int {
	numCol, numRow := len(grid), len(grid[0])
	colMap := make([]int, numCol)
	rowMap := make([]int, numRow)

	for i := 0; i < numCol; i++ {
		for j := 0; j < numRow; j++ {
			if grid[i][j] == 1 {
				colMap[i]++
				rowMap[j]++
			} else {
				colMap[i]--
				rowMap[j]--
			}
		}
	}

	for i := 0; i < numCol; i++ {
		for j := 0; j < numRow; j++ {
			grid[i][j] = colMap[i] + rowMap[j]
		}
	}

	return grid
}
