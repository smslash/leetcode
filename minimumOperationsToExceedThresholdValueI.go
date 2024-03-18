package leetcode

func minOperations(nums []int, k int) int {
    sort.Ints(nums)
    count := 0

    for i := 0; i < len(nums); i++ {
        if nums[i] < k {
            count++
        } else {
            break
        }
    }

    return count
}
