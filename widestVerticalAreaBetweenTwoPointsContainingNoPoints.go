package leetcode

func maxWidthOfVerticalArea(points [][]int) int {
	res := make([]int, len(points))
	for i := 0; i < len(points); i++ {
		res[i] = points[i][0]
	}

	sort.Ints(res)
	ans := 0

	for i := 1; i < len(res); i++ {
		n := res[i] - res[i-1]
		if n < 0 {
			n = -n
		}

		if n > ans {
			ans = n
		}
	}

	return ans
}
