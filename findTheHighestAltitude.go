package leetcode

func largestAltitude(gain []int) int {
    alt := make([]int, len(gain) + 1)
    alt[0] = 0
    res := 0
    
    for i := 1; i <= len(gain); i++ {
        res = alt[i - 1] + gain[i - 1]
        alt[i] = res
    }

    max := -100000
    for i := 0; i < len(alt); i++ {
        if max < alt[i] {
            max = alt[i]
        }
    }
    return max
}