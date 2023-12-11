package leetcode

func findSpecialInteger(arr []int) int {
    num := make(map[int]int)
    quar := len(arr) / 4 + 1
    for i := 0; i < len(arr); i++ {
        num[arr[i]]++
        if num[arr[i]] == quar {
            return arr[i]
        }
    }
    return arr[0]
}
