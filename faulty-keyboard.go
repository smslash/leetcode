package leetcode

func finalString(s string) string {
    res := ""
    for i := 0; i < len(s); i++ {
        if s[i] == 'i' {
            res = reverse(res)
        } else {
            res += string(s[i])
        }
    }
    
    return res
}

func reverse(s string) string {
	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}
