package leetcode

func maxSubarrayLength(nums []int, k int) int {
	large := 0
	l := 0
	r := 0
	table := make(map[int]int)

	if len(nums) == 1 {
		return 1
	}

	for r = 0; r < len(nums); r++ {
		table[nums[r]]++
		if table[nums[r]] <= k {
			large = max(large, r-l+1)
		}

		for table[nums[r]] > k && l < r {
			table[nums[l]]--
			l++
		}
	}

	return large
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
