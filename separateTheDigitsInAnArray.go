package leetcode

func separateDigits(nums []int) []int {
    res := make([]int, 0)
    for i := 0; i < len(nums); i++ {
        if nums[i] < 10 {
            res = append(res, nums[i])
        } else {
            tmp := make([]int, 0)
            for nums[i] > 9 {
                tmp = append(tmp, nums[i] % 10)
                nums[i] /= 10
            }
            tmp = append(tmp, nums[i])
            
            for i := len(tmp) - 1; i >= 0; i-- {
                res = append(res, tmp[i])
            }
        }
    }
    return res
}
