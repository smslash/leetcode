package leetcode

func addStrings(num1 string, num2 string) string {
    var result string

    i, j, carry := len(num1) - 1, len(num2) - 1, 0
    for i >= 0 || j >= 0 || carry > 0 {
        digit1, digit2 := 0, 0
        if i >= 0 {
            digit1 = int(num1[i] - '0')
            i--
        }
        if j >= 0 {
            digit2 = int(num2[j] - '0')
            j--
        }

        sum := digit1 + digit2 + carry
        carry = sum / 10
        sum %= 10

        result = strconv.Itoa(sum) + result
    }

    return result
}
