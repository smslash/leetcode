package leetcode

func romanToInt(s string) int {
	var sum int
	mapping := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	if len(s) == 1 {
		return mapping[s]
	}

	for i := 0; i < len(s)-1; i++ {
		current := string(s[i])
		next := string(s[i+1])
		if mapping[current] >= mapping[next] {
			sum = sum + mapping[current]
		} else {
			sum = sum - mapping[current]
		}
	}

	lastLetter := string(s[len(s)-1])
	lastValue := mapping[lastLetter]
	sum = sum + lastValue

	return sum
}
