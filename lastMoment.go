package leetcode

func getLastMoment(n int, left []int, right []int) int {
    ans := 0
    for _, v := range left {
        ans = max(ans, v)
    }
    
    for _, v := range right {
        ans = max(ans, n-v)
    }
    
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    
    return b
}
