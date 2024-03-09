package leetcode

func getCommon(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	m := make(map[int]bool)

	for i := 0; i < len(nums2); i++ {
		m[nums2[i]] = true
	}

	for i := 0; i < len(nums1); i++ {
		if m[nums1[i]] {
			return nums1[i]
		}
	}

	return -1
}
