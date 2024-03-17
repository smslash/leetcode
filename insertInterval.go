package leetcode

func insert(intervals [][]int, newInterval []int) [][]int {
	startIdx, endIdx := -1, -1
	for idx := 0; idx < len(intervals); idx++ {
		data := intervals[idx]
		if newInterval[0] <= data[1] && startIdx == -1 {
			startIdx = idx
			break
		}
	}

	for idx := len(intervals) - 1; idx >= 0; idx-- {
		data := intervals[idx]
		if newInterval[1] >= data[0] && endIdx == -1 {
			endIdx = idx
			break
		}
	}

	var result [][]int
	switch {
	case startIdx == -1:
		result = append(intervals, newInterval)
	case endIdx == -1:
		result = append([][]int{}, newInterval)
		result = append(result, intervals...)
	case startIdx > endIdx:
		result = append([][]int{}, intervals[:endIdx+1]...)
		result = append(result, newInterval)
		result = append(result, intervals[endIdx+1:]...)
	case startIdx == endIdx:
		intervals[startIdx][0] = min(intervals[startIdx][0], newInterval[0])
		intervals[startIdx][1] = max(intervals[startIdx][1], newInterval[1])
		result = intervals
	case startIdx <= endIdx:
		newData := make([]int, 2)
		newData[0] = min(intervals[startIdx][0], newInterval[0])
		newData[1] = max(intervals[endIdx][1], newInterval[1])

		result = append([][]int{}, intervals[:startIdx]...)
		result = append(result, newData)
		result = append(result, intervals[endIdx+1:]...)
	}

	return result
}
