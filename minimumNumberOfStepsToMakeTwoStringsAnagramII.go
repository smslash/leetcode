package leetcode

func minSteps(s string, t string) int {
	count, arr := 0, make([]int, 26)
	
	for _, v := range s {
		arr[v - 'a']++
	}
	for _, v := range t {
		arr[v - 'a']--
	}
	
	for _, v := range arr {
        if v < 0 {
            count += -v
        } else {
            count += v
        }
	}
	
	return count
}
