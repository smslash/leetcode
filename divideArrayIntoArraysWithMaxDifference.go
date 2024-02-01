package leetcode

func divideArray(nums []int, k int) [][]int {
	if len(nums)%3 != 0 {
		return [][]int{}
	}

	sort.Ints(nums)

	res := make([][]int, len(nums)/3)
	idx := 0
	for i := 0; i < len(nums); i += 3 {
		res[idx] = append(res[idx], nums[i])
		res[idx] = append(res[idx], nums[i+1])
		res[idx] = append(res[idx], nums[i+2])

		if nums[i+2]-nums[i] > k || nums[i+1]-nums[i] > k || nums[i+2]-nums[i+1] > k {
			return [][]int{}
		}

		idx++
	}

	return res
}
