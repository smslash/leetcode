package leetcode

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := range nums {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			switch {
			case sum > 0:
				k--
			case sum < 0:
				j++
			case sum == 0:
				result = append(result, []int{nums[i], nums[j], nums[k]})
				for j = j + 1; j < k && nums[j] == nums[j-1]; {
					j++
				}
			}
		}
	}
	return result
}
