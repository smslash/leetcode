package leetcode

func maxLength(arr []string) int {
    maxLen := 0
    backtrack(arr, 0, "", &maxLen)
    return maxLen
}

func backtrack(arr []string, pos int, current string, maxLen *int) {
    if isUnique(current) {
        if len(current) > *maxLen {
            *maxLen = len(current)
        }
    } else {
        return
    }

    for i := pos; i < len(arr); i++ {
        backtrack(arr, i+1, current+arr[i], maxLen)
    }
}

func isUnique(s string) bool {
    charSet := make(map[rune]bool)
    for _, ch := range s {
        if charSet[ch] {
            return false
        }
        charSet[ch] = true
    }
    return true
}
