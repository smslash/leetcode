package leetcode

func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	prefix := make([]int, n)
	prefix[0] = nums[0]
	for i := 1; i < n; i++ {
		prefix[i] += prefix[i-1] + nums[i]
	}

	ans := make([]int, n)

	for i := 0; i < n; i++ {
		ans[i] = i*nums[i] - (prefix[i] - nums[i]) + (prefix[n-1] - prefix[i]) - (n-1-i)*nums[i]
	}

	return ans
}
