package leetcode

func getLengthOfOptimalCompression(s string, k int) int {
	length := len(s)
	dp := make([][]int, length+1)
	for i := 0; i <= length; i++ {
		dp[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = length
		}
	}
	dp[0][0] = 0

	for i := 1; i <= length; i++ {
		for j := 0; j <= k; j++ {
			if j > 0 {
				dp[i][j] = minimum(dp[i][j], dp[i-1][j-1])
			}
			countSame, countDiff := 0, 0
			for backwardPos := i; backwardPos >= 1; backwardPos-- {
				if s[backwardPos-1] == s[i-1] {
					countSame++
				} else {
					countDiff++
				}

				if countDiff > j {
					break
				}

				dp[i][j] = minimum(dp[i][j], dp[backwardPos-1][j-countDiff]+calcCompressionLength(countSame))
			}
		}
	}
	return dp[length][k]
}

func calcCompressionLength(length int) int {
	switch {
	case length == 1:
		return 1
	case length < 10:
		return 2
	case length < 100:
		return 3
	default:
		return 4
	}
}

func minimum(val1, val2 int) int {
	if val1 < val2 {
		return val1
	}
	return val2
}
