package leetcode

func findDuplicates(nums []int) []int {
    var ans []int
    i := 0
    for i < len(nums) {

        if nums[i] <= 0 {
            i++
        } else if nums[nums[i]-1] == 0 {
            ans = append(ans,nums[i])
            nums[i] = -1
        } else {
            nums[i],nums[nums[i]-1] = nums[nums[i]-1],0
        }
    }
    return ans
}
