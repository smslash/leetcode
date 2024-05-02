package leetcode

func findMaxK(nums []int) int {
	sort.Ints(nums)

	for i := len(nums) - 1; i >= 0; i-- {
		for j := 0; j < len(nums); j++ {
			if nums[i] == -nums[j] {
				return nums[i]
			}
		}
	}

	return -1
}
