package leetcode

func sumOfSquares(nums []int) int {
    n := len(nums)
    var res int

    for i := 0; i < n; i++ {
        if n % (i + 1) == 0 {
            res += nums[i] * nums[i]
        }
    }
    
    return res
}