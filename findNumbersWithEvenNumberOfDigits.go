package leetcode

func findNumbers(nums []int) int {
    res := 0
    for i := 0; i < len(nums); i++ {
        num := 0
        for nums[i] > 0 {
            num++
            nums[i] /= 10
        }

        if num % 2 == 0 {
            res++
        }
    }
    return res
}
