package leetcode

func lengthOfLastWord(s string) int {
    if len(s) < 1 {
        return 0
    }
    
    count := 0
    for i := len(s)-1; i > -1; i-- {
        if string(s[i]) != " " {
            count++
        } else if count > 0 {
            return count 
        }
    }
    
    return count
}
