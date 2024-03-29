package leetcode

func countSubarrays(nums []int, k int) int64 {
	max := 0
	for i := range nums {
		if max < nums[i] {
			max = nums[i]
		}
	}

	var res, left = 0, 0
	count := 0
	for rigth := range nums {
		if nums[rigth] == max {
			count++
		}
		for count >= k {
			if nums[left] == max {
				count--
			}
			left++
		}
		res += left
	}
	return int64(res)
}
