package leetcode

func firstMissingPositive(nums []int) int {
    sort.Ints(nums)
    p := 1
    for i := range nums {
        if nums[i] == p { 
            p++
        }
    }
    return p
}
