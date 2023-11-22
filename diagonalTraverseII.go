package leetcode

func findDiagonalOrder(nums [][]int) []int {
	groups := make(map[int][]int)

	for i := len(nums) - 1; i >= 0; i-- {
		for j := 0; j < len(nums[i]); j++ {
			groups[i+j] = append(groups[i+j], nums[i][j])
		}
	}

	var ans []int
	var curr int

	for {
		if group, ok := groups[curr]; ok {
			ans = append(ans, group...)
			curr++
		} else {
			break
		}
	}

	return ans
}
