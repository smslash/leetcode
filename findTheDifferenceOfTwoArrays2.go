package leetcode

func findDifference(nums1 []int, nums2 []int) [][]int {
	map1 := make(map[int]bool)
	map2 := make(map[int]bool)
	res := make([][]int, 2)

	for i := 0; i < len(nums1); i++ {
		map1[nums1[i]] = true
	}

	for i := 0; i < len(nums2); i++ {
		map2[nums2[i]] = true
	}

	for i := range map1 {
		if !map2[i] {
			res[0] = append(res[0], i)
		}
	}

	for i := range map2 {
		if !map1[i] {
			res[1] = append(res[1], i)
		}
	}

	
	return res
}
