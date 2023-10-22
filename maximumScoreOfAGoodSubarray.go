func maximumScore(nums []int, k int) int {
    l, r := k, k
    res, curMin := nums[k], nums[k]
    for l > 0 || r < len(nums) - 1 {
        left, right := 0, 0
        if l > 0 {
            left = nums[l - 1]
        }
        if r < len(nums) - 1 {
            right = nums[r + 1]
        }
        if left > right {
            l--
            curMin = min(curMin, left)
            
        } else {
            r++
            curMin = min(curMin, right)
        }
        res = max(res, curMin * (r - l + 1))
    }
    return res
}

func min(a, b int) int {
    if a < b { return a }
    return b
}

func max(a, b int) int {
    if a > b { return a }
    return b
}
