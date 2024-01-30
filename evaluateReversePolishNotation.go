package leetcode

func evalRPN(tokens []string) int {
    var arith func(v1 int, v2 int, op string) int
    arith = func(v1 int, v2 int, op string) int {
        v := 0
        if op == "+" {
            v = v1 + v2
        } else if op == "-" {
            v = v1 - v2
        } else if op == "*" {
            v = v1 * v2
        } else if op == "/" {
            v = v1 / v2
        }
        return v
    }

    st := []int{}

    for _, v := range tokens {
        if v == "+" || v == "-" || v == "*" || v == "/" {
            v2 := st[len(st)-1]
            v1 := st[len(st)-2]
            st = st[:len(st)-2]
            val := arith(v1, v2, v)
            st = append(st, val)
        } else {
            n, _ := strconv.Atoi(v)
            st = append(st, n)
        }
    }

    return st[0]
}
