package leetcode

func increasingTriplet(nums []int) bool {
	first := 9223372036854775807
	second := 9223372036854775807

	for _, num := range nums {
		if num > first && num > second {
			return true
		}

		if num > first {
			second = num
		} else {
			first = num
		}
	}

	return false
}
