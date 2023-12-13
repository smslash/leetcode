package leetcode

func findDifference(nums1 []int, nums2 []int) [][]int {
	res := make([][]int, 2)
	for i := 0; i < len(nums1); i++ {
		exist := false
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				exist = true
				break
			}
		}
		if !exist {
			for k := 0; k < len(res[0]); k++ {
				if nums1[i] == res[0][k] && i != k {
					exist = true
					break
				}
			}

			if !exist {
				res[0] = append(res[0], nums1[i])
			}
		}
	}

	for i := 0; i < len(nums2); i++ {
		exist := false
		for j := 0; j < len(nums1); j++ {
			if nums2[i] == nums1[j] {
				exist = true
				break
			}
		}
		if !exist {
			for k := 0; k < len(res[1]); k++ {
				if nums2[i] == res[1][k] && i != k {
					exist = true
					break
				}
			}

			if !exist {
				res[1] = append(res[1], nums2[i])
			}
		}
	}
	return res
}
