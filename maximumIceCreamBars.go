package leetcode

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	res, count := 0, 0

	for _, v := range costs {
		res += v
		if res > coins {
			return count
		}
		count++
	}
	return count
}
