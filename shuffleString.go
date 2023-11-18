package leetcode

func restoreString(s string, indices []int) string {
    arr := make([]byte, len(s))
    for i := 0; i < len(s); i++ {
        arr[i] = s[i]
    }
    
    n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if indices[j] > indices[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
                indices[j], indices[j+1] = indices[j+1], indices[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
    
    return fmt.Sprintf("%s", arr)
}

