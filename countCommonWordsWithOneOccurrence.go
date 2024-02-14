package leetcode

func countWords(words1 []string, words2 []string) int {
    m := make(map[string]int)
    n := make(map[string]int)
    res := 0

    for i := 0; i < len(words1); i++ {
        m[words1[i]]++
    }

    for i := 0; i < len(words2); i++ {
        n[words2[i]]++
    }

    for i, v := range m {
        if value, ok := n[i]; ok {
            if v == value && v == 1 {
                res++
            }
        }
    }

    return res
}
