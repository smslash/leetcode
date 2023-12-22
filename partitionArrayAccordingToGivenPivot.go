package leetcode

func pivotArray(nums []int, pivot int) []int {
	res := []int{}

	for i := 0; i < len(nums); i++ {
		if nums[i] < pivot {
			res = append(res, nums[i])
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == pivot {
			res = append(res, nums[i])
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > pivot {
			res = append(res, nums[i])
		}
	}

	return res
}
