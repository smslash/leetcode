package leetcode

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies))
	max := candies[0]
	for i := 1; i < len(candies); i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}

	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= max {
			res[i] = true
		}
	}

	return res
}
