package leetcode

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	area := 0

	for l != r {
		if area < (r - l) * min(height[r], height[l]) {
			area = (r - l) * min(height[r], height[l])
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return area
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
