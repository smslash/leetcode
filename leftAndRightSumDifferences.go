package leetcode

func leftRightDifference(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		leftSum := 0
		for j := 0; j < i; j++ {
			leftSum += nums[j]
		}
		rightSum := 0
		for j := i + 1; j < len(nums); j++ {
			rightSum += nums[j]
		}
		n := leftSum - rightSum
		if n < 0 {
			res[i] = -n
		} else {
			res[i] = n
		}
	}
	return res
}
