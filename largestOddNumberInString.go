package leetcode

func largestOddNumber(num string) string {
	for i := len(num) - 1; i > -1; i-- {
        if isOdd(num[i]) {
            return num
        } else {
            num = num[:len(num)-1]
        }
    }
	return num
}

func isOdd(s byte) bool {
    if s == '1' || s == '3' || s == '5' || s == '7' || s == '9' {
        return true
    }
    return false
}
