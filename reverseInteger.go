func reverse(x int) int {

	ans := 0

	for ; x != 0; x /= 10 {
		ans = ans*10 + x%10
	}

	if ans > -2147483647 && ans < 2147483646 {
		return ans
	}
	return 0

}
