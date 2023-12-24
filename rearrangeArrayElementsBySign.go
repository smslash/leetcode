package leetcode

func rearrangeArray(nums []int) []int {
	res := make([]int, len(nums))
	tmp := make([]int, len(nums)/2)

	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			tmp[index] = nums[i]
			index++
		}
	}

	for i := 0; i < len(nums); i += 2 {
		res[i] = tmp[i/2]
	}

	index = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			tmp[index] = nums[i]
			index++
		}
	}

	for i := 1; i < len(nums); i += 2 {
		res[i] = tmp[i/2]
	}

	return res
}
