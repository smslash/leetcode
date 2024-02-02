package leetcode

func rotate(nums []int, k int) {
    k = k % len(nums)
    if k == 0 {
        return
    }

    temp := make([]int, 0, len(nums))
    temp = append(temp, nums[len(nums) - k:]...)
    temp = append(temp, nums[:len(nums) - k]...)

    copy(nums, temp) 
}
