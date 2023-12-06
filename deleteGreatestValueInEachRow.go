package leetcode

func deleteGreatestValue(grid [][]int) int {
    n := len(grid[0])
    ans := 0

    for n >= 0 {
        MAX := []int{}
        for i := 0 ; i < len(grid); i++ {
            max := -1
            index := -1

            for j := 0; j < len(grid[i]); j++ {
                if max < grid[i][j] {
                    max = grid[i][j]
                    index = j
                }
            }

            MAX = append(MAX, max)
            if index > -1 {
                grid[i][index] = -1
            }
        }
        var g int
        for i := 0; i < len(MAX); i++ {
            if MAX[i] > g {
                g = MAX[i]
            }
        }

        ans += g

        n--
    }

    return ans
}
