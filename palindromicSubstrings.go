package leetcode

func countSubstrings(s string) int {
	n := len(s)
	ans := 0
	dp := make([][]bool, n)

	for i := range dp {
		dp[i] = make([]bool, n)
		for j := range dp[i] {
			dp[i][j] = false
		}
	}

	for i := 0; i < n; i++ {
		dp[i][i] = true
		ans += 1
	}

	for i := 0; i < n-1; i++ {
		dp[i][i+1] = (s[i] == s[i+1])

		if dp[i][i+1] {
			ans += 1
		}
	}

	for x := 3; x < n+1; x++ {
		for i := 0; i < n-x+1; i++ {
			j := x + i - 1
			dp[i][j] = dp[i+1][j-1] && (s[i] == s[j])
			if dp[i][j] {
				ans += 1
			}
		}
	}

	return ans
}
