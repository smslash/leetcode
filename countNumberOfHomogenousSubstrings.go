package leetcode

func countHomogenous(s string) int {
    const mod = 1000000007
    total, count := 0, 1

    for i := 1; i < len(s); i++ {
        if s[i] == s[i-1] {
            count++
        } else {
            total = (total + count*(count+1)/2) % mod
            count = 1
        }
    }

    total = (total + count*(count+1)/2) % mod

    return total
}
