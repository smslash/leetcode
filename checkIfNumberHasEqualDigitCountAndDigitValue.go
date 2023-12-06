package leetcode

import "strconv"

func digitCount(num string) bool {
    for i := 0; i < len(num); i++ {
        if !check(num, i) {
            return false
        }
    }
    return true
}

func check( s string, n int) bool {
    count := 0
    for i := 0; i < len(s); i++ {
        x, _ := strconv.Atoi(string(s[i]))
        if x == n {
            count++
        }
    }

    x, _ := strconv.Atoi(string(s[n]))
    if x == count {
        return true
    }

    return false
}
