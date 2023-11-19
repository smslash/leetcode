package leetcode

func reductionOperations(nums []int) int {
	sort.Ints(nums)
	operations := 0

	i := len(nums) - 1
	for i >= 0 && nums[i] != nums[0] {
		n := nums[i]
		for nums[i] == n && i > 0 {
			i--
		}

		operations = len(nums) - 1 - i + operations
	}
	return operations
}
