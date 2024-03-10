package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)
	res := make([]int, 0)

	for i := 0; i < len(nums1); i++ {
		m[nums1[i]] = true
	}

	for i := 0; i < len(nums2); i++ {
		if m[nums2[i]] {
			res = append(res, nums2[i])
			m[nums2[i]] = false
		}
	}

	return res
}
