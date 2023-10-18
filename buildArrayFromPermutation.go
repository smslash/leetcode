package leetcode

func buildArray(nums []int) []int {
	n := len(nums)
	for idx := range nums {
		nums[idx] = nums[idx] + n*(nums[nums[idx]]%n)
	}
	for idx := range nums {
		nums[idx] = nums[idx] / n
	}
	return nums
}
