package leetcode

func removeTrailingZeros(num string) string {
    count := 0
    
    for i := len(num) - 1; i >= 0; i-- {
        if num[i] != '0' {
            break
        }
        count++
    }

    return num[:len(num)-count]
}
