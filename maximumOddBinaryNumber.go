package leetcode

func maximumOddBinaryNumber(s string) string {
	ones := ""
	zeros := ""

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			ones += "1"
		} else {
			zeros += "0"
		}
	}
	ones = ones[:len(ones)-1]

	return ones + zeros + "1"
}
