package leetcode

func findMatrix(nums []int) [][]int {
	n := make(map[int]int)
	max := 0
	for i := 0; i < len(nums); i++ {
		n[nums[i]]++
	}

	for _, v := range n {
		if v > max {
			max = v
		}
	}

	res := make([][]int, max)
	for j := 0; j < max; j++ {
		for i, v := range n {
			if v > 0 {
				res[v-1] = append(res[v-1], i)
				n[i]--
			}
		}
	}

	return res
}
