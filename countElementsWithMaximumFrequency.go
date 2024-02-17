package leetcode

func maxFrequencyElements(nums []int) int {
    m := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        m[nums[i]]++
    }

    max := -1
    for _, v := range m {
        if v > max {
            max = v
        }
    }

    count := 0
    for _, v := range m {
        if v == max {
            count++
        }
    }

    return count * max
}
