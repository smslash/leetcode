package leetcode

func minPairSum(nums []int) int {
	sort.Ints(nums)
	i, j := 0, len(nums)-1
	max := nums[i] + nums[j]
	var tmp int
	i += 1
	j -= 1
	for i < j {
		tmp = nums[i] + nums[j]
		if tmp > max {
			max = tmp
		}
		i, j = i+1, j-1
	}
	return max
}
