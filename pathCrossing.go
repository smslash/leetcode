package leeetcode

func isPathCrossing(path string) bool {
	x := make(map[[2]int]int)
	point := [2]int{0, 0}
	x[point]++

	for _, v := range path {
		switch v {
		case 'N':
			point[0] += 1
		case 'S':
			point[0] -= 1
		case 'E':
			point[1] += 1
		case 'W':
			point[1] -= 1
		}

		x[point]++
        
		if x[point] > 1 {
			return true
		}
	}

	return false
}
