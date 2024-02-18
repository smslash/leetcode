package leetcode

func lessOrEqual(candles []int, idx int) int {
	if len(candles) == 0 {
		return -1
	}
	l, r := 0, len(candles)-1
	for l+1 < r {
		mid := l + (r-l)/2
		if candles[mid] < idx {
			l = mid
		} else {
			r = mid
		}
	}
	if candles[r] <= idx {
		return candles[r]
	}
	if candles[l] <= idx {
		return candles[l]
	}
	return -1
}

func upperOrEqual(candles []int, idx int) int {
	if len(candles) == 0 {
		return -1
	}

	l, r := 0, len(candles)-1
	for l+1 < r {
		mid := l + (r-l)/2
		if candles[mid] <= idx {
			l = mid
		} else {
			r = mid
		}
	}

	if candles[l] >= idx {
		return candles[l]
	}
	if candles[r] >= idx {
		return candles[r]
	}
	return -1
}

func platesBetweenCandles(s string, queries [][]int) []int {
	prefixPlates := make([]int, len(s)+1)
	prefixPlates[0] = 0
	candles := make([]int, 0, len(s))
	for i := 1; i < len(prefixPlates); i++ {
		prefixPlates[i] = prefixPlates[i-1]
		if s[i-1] == '*' {
			prefixPlates[i]++
		} else {
			candles = append(candles, i-1)
		}
	}

	out := make([]int, 0, len(queries))
	for _, query := range queries {
		left := upperOrEqual(candles, query[0])
		right := lessOrEqual(candles, query[1])
		if right == -1 || left == -1 || right <= left {
			out = append(out, 0)
		} else {
			out = append(out, prefixPlates[right]-prefixPlates[left])
		}
	}
	return out
}
