package leetcode

func numSubarrayProductLessThanK(nums []int, k int) int {
	count := 0

	for i := 0; i < len(nums); i++ {
		sum := 1
		for j := i; j < len(nums); j++ {
			sum *= nums[j]
			if sum < k {
				count++
			} else {
				break
			}
		}
	}

	return count
}
