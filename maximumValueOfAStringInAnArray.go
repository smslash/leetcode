package leetcode

func isNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func maximumValue(strs []string) int {
	var max int
	for _, str := range strs {
		var value int
		if isNumeric(str) {
			value, _ = strconv.Atoi(str)
		} else {
			value = len(str)
		}
		if value > max {
			max = value
		}
	}
	return max
}
