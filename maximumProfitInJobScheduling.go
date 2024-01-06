package leetcode

type Job struct {
    start  int
    end    int
    profit int
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
    n := len(startTime)
    jobs := make([]Job, n)
    
    for i := 0; i < n; i++ {
        jobs[i] = Job{startTime[i], endTime[i], profit[i]}
    }
    
    sort.Slice(jobs, func(i, j int) bool {
        return jobs[i].end < jobs[j].end
    })
    
    dp := make([]int, n)
    dp[0] = jobs[0].profit
    
    for i := 1; i < n; i++ {
        currentProfit := jobs[i].profit
        for j := i - 1; j > -1; j-- {
            if jobs[j].end <= jobs[i].start {
                currentProfit += dp[j]
                break
            }
        }
        dp[i] = max(dp[i-1], currentProfit)
    }
    
    return dp[n-1]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
