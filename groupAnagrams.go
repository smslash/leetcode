package leetcode

func groupAnagrams(strs []string) [][]string {
	var res [][]string
	var equal bool
	check := make([]bool, len(strs))

	for i := 0; i < len(strs); i++ {
		if !check[i] {
			res = append(res, []string{strs[i]})
			for j := i + 1; j < len(strs); j++ {
				if !check[j] {
					if len(strs[i]) == len(strs[j]) {
						equal = true
						map1 := initMap(strs[i])
						map2 := initMap(strs[j])

						for key, value1 := range map1 {
							value2, ok := map2[key]
							if !ok || value1 != value2 {
								equal = false
							}
						}

						if equal {
							check[i], check[j] = true, true
							res[len(res)-1] = append(res[len(res)-1], strs[j])
						}
					}
				}
			}
		}
	}

	return res
}

func initMap(s string) map[byte]int {
	a := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		a[s[i]]++
	}

	return a
}
