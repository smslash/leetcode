package leetcode

func maxFrequency(nums []int, k int) int {
    sort.Ints(nums)
    left, ans, curr := 0, 0, 0

    for right := 0; right < len(nums); right++ {
        target := nums[right]
        curr += target

        for (right - left + 1) * target - curr > k {
            curr -= nums[left]
            left++
        }

        ans = max(ans, right - left + 1)
    }

    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
