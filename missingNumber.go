package leetcode

func missingNumber(nums []int) int {
    a, b := len(nums), 0
    for i := 0; i < len(nums); i++ {
        a += i
        b += nums[i]
    }

    return a - b
}
