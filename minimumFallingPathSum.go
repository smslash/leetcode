package leetcode

func minFallingPathSum(matrix [][]int) int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return 0
    }

    for row := len(matrix) - 2; row >= 0; row-- {
        for col := 0; col < len(matrix[0]); col++ {
            minNext := matrix[row+1][col]
            if col > 0 {
                minNext = min(minNext, matrix[row+1][col-1])
            }
            if col < len(matrix[0])-1 {
                minNext = min(minNext, matrix[row+1][col+1])
            }

            matrix[row][col] += minNext
        }
    }

    minSum := matrix[0][0]
    for _, val := range matrix[0] {
        minSum = min(minSum, val)
    }
    return minSum
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
