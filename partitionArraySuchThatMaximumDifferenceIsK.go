package leetcode

func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	ans := 1
	min := nums[0]
	max := nums[0]

	for i := 0; i < len(nums); i++ {
		max = nums[i]

		if (max - min) > k {
			ans++

			min = nums[i]
		}
	}

	return ans
}
