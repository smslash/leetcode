package leetcode

func mySqrt(x int) int {
	for i := 0; i <= x; i++ {
		if i*i <= x && (i+1)*(i+1) > x {
			return i
		}
	}
	return 0
}
