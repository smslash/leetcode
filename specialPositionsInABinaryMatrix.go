package leetcode

func numSpecial(mat [][]int) int {
	res := 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				breaker := false
				for row := 0; row < len(mat); row++ {
					if row == i {
						continue
					}

					if mat[row][j] == 1 {
						breaker = true
						break
					}
				}

				if breaker {
					continue
				}

				for column := 0; column < len(mat[i]); column++ {
					if column == j {
						continue
					}

					if mat[i][column] == 1 {
						breaker = true
						break
					}
				}

				if breaker {
					continue
				}

				res++
			}
		}
	}
	return res
}
