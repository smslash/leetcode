package leetcode

func countKDifference(nums []int, k int) int {
    count := 0
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            tmp := nums[i] - nums[j]
            if tmp < 0 {
                tmp = -tmp
                if tmp == k {
                    count++
                }
            } else if tmp == k {
                count++
            }
        }
    }
    return count
}
