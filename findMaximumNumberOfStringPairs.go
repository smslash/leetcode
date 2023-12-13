package leetcode

func maximumNumberOfStringPairs(words []string) int {
    res := 0
    for i := 0; i < len(words); i++ {
        for j := 0; j < len(words); j++ {
            if i == j {
                continue
            }

            if words[i] == reverseString(words[j]) {
                res++
            }
        }
    }
    return res / 2
}

func reverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
