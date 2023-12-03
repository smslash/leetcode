package leetcode

func minTimeToVisitAllPoints(points [][]int) int {
	steps := 0

	for i := 0; i < len(points) - 1; i++ {
		x := points[i][0] - points[i+1][0]
		y := points[i][1] - points[i+1][1]
		
		if x < 0 {
			x = -x
		}

		if y < 0 {
			y = -y
		}

		if x < y {
			steps += y
		} else {
			steps += x
		}
	}

	return steps
}
