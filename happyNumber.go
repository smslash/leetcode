package leetcode

func isHappy(n int) bool {
    m := make(map[int]bool)
    m[n] = true

    for n != 1 {
        sum := 0
        for n != 0 {
            num := n%10
            sum += num * num
            n /= 10
        }

        n = sum

        if m[n] {
            return false
        }

        m[n] = true
    }

    return true
}
