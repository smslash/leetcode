package leetcode

func largestGoodInteger(num string) string {
	max := -1
	var res string

	for i := 0; i < len(num)-2; i++ {
		if num[i] == num[i+1] && num[i] == num[i+2] {
			str := string(num[i]) + string(num[i+1]) + string(num[i+2])
			n, _ := strconv.Atoi(str)
			if max < n {
				max = n
				res = str
			}
		}
	}
	return res
}
