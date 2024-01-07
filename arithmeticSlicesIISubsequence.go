package leetcode

func numberOfArithmeticSlices(nums []int) int {
    count := 0
    if len(nums) < 3 {
        return count
    }
    
    dp := make([]map[int]int, len(nums))
    
    for i := range nums {
        dp[i] = make(map[int]int)
        for j := 0; j < i; j++ {
            diff := nums[i] - nums[j]
            countJ, foundJ := dp[j][diff]
            countI, foundI := dp[i][diff]
            if !foundI {
                dp[i][diff] = 0
            }
            if !foundJ {
                countJ = 0
            }
            dp[i][diff] = countI + countJ + 1
            count += countJ
        }
    }
    
    return count
}
