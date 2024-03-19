package leetcode

func leastInterval(tasks []byte, n int) int {
	f := 0
	m := make(map[byte]int)

	for i := 0; i < len(tasks); i++ {
		m[tasks[i]]++
		if m[tasks[i]] > f {
			f = m[tasks[i]]
		}
	}

	res := (f - 1) * (n + 1)
	for _, value := range m {
		if value == f {
			res++
		}
	}

	return max(res, len(tasks))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
