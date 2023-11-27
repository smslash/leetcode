package leetcode

func majorityElement(nums []int) int {
    count := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        count[nums[i]]++
    }

    max := -1000000000
    var ans int
    for i, v := range count {
        if v > max {
            max = v
            ans = i
        }
    }

    return ans
}
