package leetcode

func sortTheStudents(score [][]int, k int) [][]int {
	for i := 0; i < len(score); i++ {
		for j := 0; j < len(score)-i-1; j++ {
			if score[j+1][k] > score[j][k] {
				score[j+1], score[j] = score[j], score[j+1]
			}
		}
	}
	return score
}
