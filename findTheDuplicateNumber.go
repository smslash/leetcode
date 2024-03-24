package leetcode

 func findDuplicate(nums []int) int {
  i := 0

  for i < len(nums) {
    num := nums[i]
    if num == 0 {
      i++
    } else if nums[num] == 0 {
      return num
    } else {
      tmp := nums[num]
      nums[num] = 0
      nums[i] = tmp
    }
  }

  return 0
}
