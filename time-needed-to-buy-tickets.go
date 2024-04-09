package leetcode

func timeRequiredToBuy(tickets []int, k int) int {
	count := 0

	for tickets[k] != 0 {
		for i := 0; i < len(tickets); i++ {
			if tickets[i] > 0 {
				count++
				tickets[i]--
			}

			if tickets[k] == 0 {
				break
			}
		}
	}

	return count
}
