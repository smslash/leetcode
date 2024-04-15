package leetcode

func distinctDifferenceArray(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		suff := make(map[int]bool)
		pref := make(map[int]bool)

		for l := 0; l <= i; l++ {
			suff[nums[l]] = true
		}

		for m := i + 1; m < len(nums); m++ {
			pref[nums[m]] = true
		}

		res[i] = len(suff) - len(pref)
	}

	return res
}
