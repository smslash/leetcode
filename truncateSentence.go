package leetcode

func truncateSentence(s string, k int) string {
    str := strings.Fields(s)
    res := ""
    for i := 0; i < k; i++ {
        res += str[i]
        if i + 1 < k {
            res += " "
        }
    }
    return res
}
