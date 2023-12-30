package leetcode

func uniqueOccurrences(arr []int) bool {
	count := make(map[int]int)
	unique := make(map[int]bool)

	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}

	for _, v := range count {
		if !unique[v] {
			unique[v] = true
		} else {
			return false
		}
	}

	return true
}
