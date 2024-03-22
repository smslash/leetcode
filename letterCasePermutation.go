package leetcode

func letterCasePermutation(s string) []string {
    sl := strings.ToLower(s)
    bs := []byte(sl)
    res := []string{}
    backtrace(&res, bs, 0)
    return res
}

func backtrace(res *[]string, bs []byte, start int) {   
    *res = append(*res, string(bs))
    for i := start; i < len(bs); i++ {
        if bs[i] >= 'a' && bs[i] <= 'z' {
            bs[i] -= 32
            backtrace(res, bs, i+1)
            bs[i] += 32
        }
    }
}
