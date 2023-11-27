package leetcode

const MOD = 1000000007

var moves = [][]int{
    {4, 6},
    {6, 8},
    {7, 9},
    {4, 8},
    {3, 9, 0},
    {},
    {1, 7, 0},
    {2, 6},
    {1, 3},
    {2, 4},
}

func knightDialer(n int) int {
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, 10)
    }

    for i := 0; i <= 9; i++ {
        dp[1][i] = 1
    }

    for i := 2; i <= n; i++ {
        for j := 0; j <= 9; j++ {
            for _, move := range moves[j] {
                dp[i][j] = (dp[i][j] + dp[i-1][move]) % MOD
            }
        }
    }

    sum := 0
    for j := 0; j <= 9; j++ {
        sum = (sum + dp[n][j]) % MOD
    }

    return sum
}
