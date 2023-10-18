package leetcode

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	sum := 0

	var s [100]int
	for i := 0; x > 0; x /= 10 {
		s[i] = x % 10
		i++
		sum++
	}
	fmt.Println(s)
	for i := 0; i < sum/2; i++ {
		if s[i] == s[sum-i-1] {
			continue
		}
		return false
	}
	return true
}
