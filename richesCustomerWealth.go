package leetcode

func maximumWealth(accounts [][]int) int {
	max := -9223372036854775808
	for i := 0; i < len(accounts); i++ {
		count := 0
		for j := 0; j < len(accounts[i]); j++ {
			count += accounts[i][j]
		}

		if max < count {
			max = count
		}
	}
	return max
}
