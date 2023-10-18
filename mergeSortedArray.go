package leetcode

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	a := 0
	for i := m; i < len(nums1); i++ {
		nums1[i] = nums2[a]
		a++
	}

	sort.Ints(nums1)
}
