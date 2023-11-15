package leetcode

func searchPosition(nums []int, left, right, target int, mode byte) int {
	res := -1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			res = mid
			if mode == 's' {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return res
}

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	lastIndex := len(nums) - 1
	res[0] = searchPosition(nums, 0, lastIndex, target, 's')
	res[1] = searchPosition(nums, 0, lastIndex, target, 'e')
	return res
}
