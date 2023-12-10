package leetcode

func transpose(matrix [][]int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == 0 {
				res = append(res, make([]int, 0))
			}
			if i == j {
				res[j] = append(res[j], matrix[i][j])
			} else {
				res[j] = append(res[j], matrix[i][j])
			}
		}
	}
	return res
}
