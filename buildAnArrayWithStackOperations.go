package leetcode

func buildArray(target []int, n int) []string {
    ans := []string{}
    i := 0
    for _, v := range target {
        for i < v - 1 {
            ans = append(ans, "Push")
            ans = append(ans, "Pop")
            i++
        }
        ans = append(ans, "Push")
        i++
    }
    return ans
}
