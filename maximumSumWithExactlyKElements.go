package leetcode

func maximizeSum(nums []int, k int) int {
    var max int
    var ans int

    for i := 0; i < len(nums); i++ {
        if max < nums[i] {
            max = nums[i]
        }
    }

    for i := 0; i < k; i++ {
        ans += max
        max++
    }

    return ans
}