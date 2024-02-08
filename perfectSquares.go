package leetcode

func numSquares(n int) int {
    m := make([]int, n + 1)
    m[0], m[1] = 0, 1

    for i := 2; i < n + 1; i++ {
        current := math.MaxInt32
        for j := 1; j * j <= i; j++ {
            current = min(current, m[i - j*j] + 1)
        }
        m[i] = current
    }
    
    return m[n]
}
