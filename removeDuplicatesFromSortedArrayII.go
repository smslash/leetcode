package leetcode

func removeDuplicates(nums []int) int {
	const big = 10001
	k, same, tmp := 0, 0, nums[0]

	for i := 0; i < len(nums); i++ {
		if tmp != nums[i] {
			tmp = nums[i]
			same = 0
		}

		if same > 1 {
			nums[i] = big
			k++
		}
		same++
	}

	sort.Ints(nums)
	return len(nums) - k
}
