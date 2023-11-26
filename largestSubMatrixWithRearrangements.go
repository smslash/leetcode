package leetcode

func largestSubmatrix(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	ans := 0

	for row := 1; row < m; row++ {
		for col := 0; col < n; col++ {
			if matrix[row][col] == 1 {
				matrix[row][col] += matrix[row-1][col]
			}
		}
	}

	for _, heights := range matrix {
		sortedHeights := make([]int, n)
		copy(sortedHeights, heights)
		sort.Slice(sortedHeights, func(i, j int) bool {
			return sortedHeights[i] > sortedHeights[j]
		})

		for i := 0; i < n; i++ {
			ans = max(ans, sortedHeights[i]*(i+1))
		}
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
