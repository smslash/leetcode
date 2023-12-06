package leetcode

func repeatedCharacter(s string) byte {
    ans := make(map[byte]int)

    for i := 0; i < len(s); i++ {
        ans[s[i]]++
        if ans[s[i]] == 2 {
            return s[i]
        }
    }
    return 0
}