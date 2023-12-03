package leetcode

func minOperations(boxes string) []int {
	res := make([]int, len(boxes))

	for i := 0; i < len(boxes); i++ {
		for j := 0; j < len(boxes); j++ {
			if boxes[j] == '0' {
				continue
			}

			if j > i {
				res[i] += j - i
			} else {
				res[i] += i - j
			}
		}
	}

	return res
}
