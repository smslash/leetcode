package leetcode

func maxProfit(prices []int) int {
	minIndex := -1
	maxIndex := -1
	min := 10000
	max := -1
	res := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
			minIndex = i
		}

		if prices[i] > max {
			max = prices[i]
			maxIndex = i
		}

		if minIndex < maxIndex {
			if max-min > res {
				res = max - min
			}
		} else {
			max = -1
			maxIndex = -1
		}
	}

	return res
}
