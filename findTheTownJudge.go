package leetcode

func findJudge(n int, trust [][]int) int {
    trustArray := make([]int, n)
    trustMap := make(map[int]bool)
    for _, val := range trust {
        trustArray[val[1]-1] += 1
        trustMap[val[0]-1] = true
    }
    for key, val := range trustArray {
        if val == n-1 && !trustMap[key]{
            return key+1
        }
    }
    return -1
}
