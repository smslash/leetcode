package leetcode

func maxSum(grid [][]int) int {
	max, sum := 0, 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i+2 < len(grid) && j+2 < len(grid[i]) {
				sum = grid[i][j] + grid[i][j+1] + grid[i][j+2] + grid[i+1][j+1] + grid[i+2][j] + grid[i+2][j+1] + grid[i+2][j+2]
			}
			if sum > max {
				max = sum
			}
		}
	}
	return max
}
