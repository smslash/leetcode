package leetcode

type Roman struct {
	S string
	V int
}

func intToRoman(num int) string {

	r := []Roman{
		{"I", 1}, {"IV", 4}, {"V", 5}, {"IX", 9}, {"X", 10}, {"XL", 40}, {"L", 50}, {"XC", 90}, {"C", 100}, {"CD", 400}, {"D", 500}, {"CM", 900}, {"M", 1000},
	}

	len := 12
	ans := ""

	for i := len; i >= 0; i-- {

		if times := num / r[i].V; times != 0 {
			if times == 1 {
				ans += r[i].S
			} else {
				for j := 1; j <= times; j++ {
					ans += r[i].S
				}
			}
			num %= r[i].V
		}

	}

	return ans
}
