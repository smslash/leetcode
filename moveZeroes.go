package leetcode

func moveZeroes(nums []int) {
    left := 0
    for right := 0; right < len(nums); right++ {
        if nums[right] != 0 {
            if left != right {
                nums[left], nums[right] = nums[right], nums[left]
            }
            left++
        }
    }
}
