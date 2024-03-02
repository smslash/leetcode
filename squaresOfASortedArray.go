package leetcode

func sortedSquares(nums []int) []int {
    n := len(nums)
    result := make([]int, n)
    n--

    for i, j := 0, n; n >= 0; n-- {     
        if nums[i] < 0 && -nums[i] >= nums[j] {
            result[n] = nums[i] * nums[i]
            i++
        } else {
            result[n] = nums[j] * nums[j]
            j--
        }
    }
    
    return result
}
