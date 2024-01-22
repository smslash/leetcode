package leetcode

func findErrorNums(nums []int) []int {
    n := len(nums)
    duplicate := -1
    sum := 0
    seen := make(map[int]bool)

    for _, num := range nums {
        if seen[num] {
            duplicate = num
        }
        seen[num] = true
        sum += num
    }

    missing := ((n * (n + 1)) / 2) - (sum - duplicate)
    return []int{duplicate, missing}
}
