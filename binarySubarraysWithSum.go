package leetcode

func numSubarraysWithSum(nums []int, goal int) int {
	count := 0
	l := len(nums)
	flag := false
	sum := 0

	for i := 0; i < l; i++ {
		sum = 0
		flag = false
		for j := i; j < l; j++ {
			sum += nums[j]
			if sum == goal {
				count++
				flag = true
			}

			if flag && sum != goal {
				flag = false
				break
			}
		}
	}

	return count
}
