package leetcode

func removeOccurrences(s string, part string) string {
    for {
        var exist bool
        for i := 0; i < len(s) - len(part) + 1; i++ {
            if s[i:len(part)+i] == part {
                s = s[:i] + s[len(part)+i:]
                exist = true
                break
            }
        }

        if !exist {
            break
        }
    }

    return s
}
