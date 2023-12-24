package leetcode

func minOperations(s string) int {
	start0 := 0
    start1 := 0
    for i := 0; i < len(s); i++ {
        if i % 2 == 0 {
            if s[i] == '0' {
                start1++
            } else {
                start0++
            }
        } else {
            if s[i] == '1' {
                start1++
            } else {
                start0++
            }
        }
    }

    if start1 < start0 {
        return start1
    }

    return start0
}
