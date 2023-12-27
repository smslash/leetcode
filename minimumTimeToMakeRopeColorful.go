package leetcode

func minCost(colors string, time []int) int {
    result := 0
    for i := 1; i < len(colors); i++ {
        if colors[i-1] == colors[i] {
            if time[i-1] < time[i] {
                result += time[i-1]
            } else {
                result += time[i]
                time[i] = time[i-1]
            }
        }
    }
    return result
}
