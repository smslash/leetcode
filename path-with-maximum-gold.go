package leetcode

func getMaximumGold(grid [][]int) int {
	var (
		m   = len(grid)
		n   = len(grid[0])
		res = 0
	)

	var dfs func(row, col int, visited [][]bool, currentGold int)
	dfs = func(row, col int, checked [][]bool, cur int) {
		if row < 0 || row >= m || col < 0 || col >= n || grid[row][col] == 0 || checked[row][col] {
			return
		}

		checked[row][col] = true
		cur += grid[row][col]
		if cur > res {
			res = cur
		}

		dfs(row+1, col, checked, cur)
		dfs(row-1, col, checked, cur)
		dfs(row, col+1, checked, cur)
		dfs(row, col-1, checked, cur)

		checked[row][col] = false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				visited := make([][]bool, m)
				for k := range visited {
					visited[k] = make([]bool, n)
				}
				dfs(i, j, visited, 0)
			}
		}
	}

	return res
}
