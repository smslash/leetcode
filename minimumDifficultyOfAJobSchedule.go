package leetcode

func minInt(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func maxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}

var cache [][]int

func dp(jobDifficulty []int, d, i int) int {
	n := len(jobDifficulty)
	if n < d || i >= n {
		return math.MaxInt32
	}
	if d == 1 {
		max := 0
		for j := i; j < n; j++ {
			max = maxInt(max, jobDifficulty[j])
		}
		return max
	}

	if cache[d][i] != -1 {
		return cache[d][i]
	}

	res := math.MaxInt32
	max := 0
	for j := i; j < n; j++ {
		max = maxInt(max, jobDifficulty[j])
		res = minInt(res, dp(jobDifficulty, d-1, j+1)+max)
	}
	cache[d][i] = res
	return res
}

func minDifficulty(jobDifficulty []int, d int) int {
	if len(jobDifficulty) < d {
		return -1
	}
	n := len(jobDifficulty)
	cache = make([][]int, d+1)
	for i := 0; i <= d; i++ {
		cache[i] = make([]int, n)
		for j := 0; j < n; j++ {
			cache[i][j] = -1
		}
	}
	res := dp(jobDifficulty, d, 0)
	if res >= math.MaxInt32 {
		return -1
	}
	return res
}
