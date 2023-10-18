package leetcode

func numIdenticalPairs(nums []int) int {
	res := make(map[int]int)
	var ans int

	for _, num := range nums {
		ans += res[num]
		res[num]++
	}
	return ans
}
