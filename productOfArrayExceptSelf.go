package leetcode

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	nx := 1
	for i := 0; i < len(nums); i++ {
		res[i] = nx
		nx *= nums[i]
	}
	nx = 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= nx
		nx *= nums[i]
	}
	return res
}
