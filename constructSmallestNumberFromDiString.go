package leetcode

func smallestNumber(pattern string) string {
    n := len(pattern) + 1
    used := make([]bool, 10)
    num := make([]byte, n)
    var solution string
    var found bool

    var buildNum func(idx, lastDigit int)
    buildNum = func(idx, lastDigit int) {
        if found {
            return
        }
        if idx == n {
            solution = string(num)
            found = true
            return
        }

        start, end, step := 1, 10, 1
        if idx > 0 && pattern[idx-1] == 'D' {
            start, end, step = lastDigit-1, 0, -1
        }

        for i := start; i != end; i += step {
            if !used[i] {
                used[i] = true
                num[idx] = '0' + byte(i)
                buildNum(idx+1, i)
                used[i] = false
            }
        }
    }

    buildNum(0, 0)
    return solution
}
