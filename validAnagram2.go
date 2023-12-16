package leetcode

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    ss := sortRunes(s)
    tt := sortRunes(t)

    for i := 0; i < len(s); i++ {
        if ss[i] != tt[i] {
            return false
        }
    }

    return true
}

func sortRunes(str string) []rune {
    runes := []rune(str)
    sort.Slice(runes, func(i, j int) bool {
        return runes[i] < runes[j]
    })
    return runes
}
