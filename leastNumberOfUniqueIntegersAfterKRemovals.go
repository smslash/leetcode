package leetcode

func findLeastNumOfUniqueInts(arr []int, k int) int {
	freqMap := make(map[int]int)
	for _, num := range arr {
		freqMap[num]++
	}

	frequencies := make([]int, 0, len(freqMap))
	for _, freq := range freqMap {
		frequencies = append(frequencies, freq)
	}

	sort.Ints(frequencies)

	for i := 0; i < len(frequencies) && k > 0; i++ {
		if frequencies[i] <= k {
			k -= frequencies[i]
			frequencies[i] = 0
		} else {
			frequencies[i] -= k
			k = 0
		}
	}

	uniqueCount := 0
	for _, freq := range frequencies {
		if freq > 0 {
			uniqueCount++
		}
	}

	return uniqueCount
}
