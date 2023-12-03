package leetcode

func executeInstructions(n int, startPos []int, s string) []int {
	answer := make([]int, len(s))
	pos := [2]int{0, 0}

	for i := 0; i < len(s); i++ {
		pos[0] = startPos[0]
		pos[1] = startPos[1]
		for j := i; j < len(s); j++ {
			if s[j] == 'R' {
				if pos[1]+1 < n {
					answer[i]++
					pos[1]++
				} else {
					break
				}
			} else if s[j] == 'L' {
				if pos[1]-1 > -1 {
					answer[i]++
					pos[1]--
				} else {
					break
				}
			} else if s[j] == 'D' {
				if pos[0]+1 < n {
					answer[i]++
					pos[0]++
				} else {
					break
				}
			} else if s[j] == 'U' {
				if pos[0]-1 > -1 {
					answer[i]++
					pos[0]--
				} else {
					break
				}
			}
		}
	}

	return answer
}
