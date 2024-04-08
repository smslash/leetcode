package leetcode

func numberOfSubarrays(nums []int, k int) int {
	tc, oc, cn := 0, 0, 0

	s := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 != 0 {
			oc++
			cn = 0
		}
		for oc == k {
			cn++
			if nums[s]%2 != 0 {
				oc--
			}
			s++
		}
		tc += cn
	}
	return tc
}
