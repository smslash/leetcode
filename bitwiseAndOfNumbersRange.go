package leetcode

func rangeBitwiseAnd(left int, right int) int {
	res := left
	for i := left + 1; i <= right; i++ {
		res = res & i
		if res == 0 {
			break
		}
	}
	return res
}
