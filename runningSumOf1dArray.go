package leetcode

func runningSum(nums []int) []int {
	current := 0
	res := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		res[i] = nums[i] + current
		current += nums[i]
	}

	return res
}
