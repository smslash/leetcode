package leetcode

func findThePrefixCommonArray(A []int, B []int) []int {
	C := make([]int, len(A))

	for i := 0; i < len(A); i++ {
		count := 0
		amap := make(map[int]bool)
		bmap := make(map[int]bool)
		for j := 0; j <= i; j++ {
			amap[A[j]] = true
			bmap[B[j]] = true
		}

		for i := range amap {
			if amap[i] == bmap[i] {
				count++
			}
		}

		C[i] = count
	}

	return C
}
