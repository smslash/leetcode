package leetcode

func getConcatenation(nums []int) []int {
	ans := make([]int, 2*len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = nums[i]
	}
	j := 0
	for i := len(nums); i < len(ans); i++ {
		ans[i] = nums[j]
		j++
	}
	return ans
}
