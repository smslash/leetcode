package leetcode

func subsetXORSum(nums []int) int {
    return bt(nums, 0)
}

func bt(nums []int, x int) int {
    if len(nums) == 0 {
        return x
    }
    
    with := bt(nums[1:], x ^ nums[0])
    without := bt(nums[1:], x)
    
    return with + without
}
