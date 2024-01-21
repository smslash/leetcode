package leetcode

func rob(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }
    if n == 1 {
        return nums[0]
    }
    
    dp := make([]int, n+1)
    dp[n], dp[n-1] = 0, nums[n-1]
    
    for i := n - 2; i >= 0; i-- {
        dp[i] = max(dp[i+1], nums[i]+dp[i+2])
    }
    
    return dp[0]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
