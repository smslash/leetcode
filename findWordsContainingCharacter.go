package leetcode

func findWordsContaining(words []string, x byte) []int {
    var arr []int
    for i := 0; i < len(words); i++ {
        for j := 0; j < len(words[i]); j++ {
            if words[i][j] == x {
                arr = append(arr, i)
                break
            }
        }
    }
    return arr
}
