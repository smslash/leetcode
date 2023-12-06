package leetcode

func totalMoney(n int) int {
    if n == 0 {
        return 0
    }

    monday, dollar, res := 1, 0, 0

    for i := 1; i <= n; i++ {
        if (i - 1) % 7 == 0 {
            dollar = monday
            monday++
        } else {
            dollar++
        }
        res += dollar
        fmt.Println(dollar)
    }
    
    return res
}
