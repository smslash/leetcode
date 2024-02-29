package leetcode

func numberGame(nums []int) []int {
    sort.Ints(nums)
    i := 0
    j := 1

    for j < len(nums) {
        nums[i], nums[j] = nums[j], nums[i]
        i, j = j + 1, j + 2
    }

    return nums
}
