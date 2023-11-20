package leetcode

func garbageCollection(garbage []string, travel []int) int {
	mTruck, pTruck, gTruck := 0, 0, 0
	mCount, pCount, gCount := 0, 0, 0
	mIndex, pIndex, gIndex := 0, 0, 0
	for i := 0; i < len(garbage); i++ {
		for j := 0; j < len(garbage[i]); j++ {
			switch garbage[i][j] {
			case 'M':
				mIndex = i
				mCount++
			case 'P':
				pIndex = i
				pCount++
			case 'G':
				gIndex = i
				gCount++
			}
		}
	}

	for i := 0; i < len(garbage); i++ {
		for j := 0; j < len(garbage[i]); j++ {
			switch garbage[i][j] {
			case 'M':
				mTruck++
			case 'P':
				pTruck++
			case 'G':
				gTruck++
			}
		}
	}

	// fmt.Println(mCount + pCount + gCount)

	res := 0
	for i := 0; i < len(travel); i++ {
		if i <= mIndex-1 && mIndex != 0 {
			res += travel[i]
			// fmt.Printf("M: %d\n", i)
		}

		if i <= pIndex-1 && pIndex != 0 {
			res += travel[i]
			// fmt.Printf("P: %d\n", i)
		}

		if i <= gIndex-1 && gIndex != 0 {
			res += travel[i]
			// fmt.Printf("G: %d\n", i)
		}
	}

	return res + mCount + pCount + gCount
}
