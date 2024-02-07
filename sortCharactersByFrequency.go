package leetcode

func frequencySort(s string) string {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	count := make([]int, 0, len(m))
	char := make([]byte, 0, len(m))

	for i, v := range m {
		count = append(count, v)
		char = append(char, i)
	}

	quickSort(count, char)

	res := make([]byte, 0, len(s))
	for i := 0; i < len(count); i++ {
		for count[i] > 0 {
			res = append(res, char[i])
			count[i]--
		}
	}

	return string(res)
}

func quickSort(slice []int, char []byte) {
    if len(slice) < 2 {
        return
    }

    left, right := 0, len(slice)-1

    pivotIndex := len(slice) / 2

    slice[pivotIndex], slice[right] = slice[right], slice[pivotIndex]
    char[pivotIndex], char[right] = char[right], char[pivotIndex]

    for i := range slice {
        if slice[i] > slice[right] {
            slice[left], slice[i] = slice[i], slice[left]
            char[left], char[i] = char[i], char[left]
            left++
        }
    }

    slice[left], slice[right] = slice[right], slice[left]
    char[left], char[right] = char[right], char[left]

    quickSort(slice[:left], char[:left])
    quickSort(slice[left+1:], char[left+1:])
}
