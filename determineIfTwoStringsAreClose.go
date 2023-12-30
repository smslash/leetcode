package leetcode

func closeStrings(word1 string, word2 string) bool {
	map1, map2 := make(map[byte]int), make(map[byte]int)

	for i := 0; i < len(word1); i++ {
		map1[word1[i]]++
	}

	for i := 0; i < len(word2); i++ {
		map2[word2[i]]++
	}

	if len(map1) != len(map2) {
		return false
	}

	arr1, arr2 := make([]int, len(map1)), make([]int, len(map1))

	i := 0
	for char, v := range map1 {
		arr1[i] = v
		if map2[char] == 0 {
			return false
		}
		i++
	}

	i = 0
	for char, v := range map2 {
		arr2[i] = v
		if map1[char] == 0 {
			return false
		}
		i++
	}

	sort.Ints(arr1)
	sort.Ints(arr2)

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}
