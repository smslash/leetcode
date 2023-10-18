package leetcode

import "strconv"

func sumIndicesWithKSetBits(nums []int, k int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		s := decimalToBinary(i)
		count := 0
		for j := 0; j < len(s); j++ {
			if s[j] == '1' {
				count++
			}
		}
		if count == k {
			res += nums[i]
		}
	}

	return res
}

func decimalToBinary(n int) string {
	res := ""
	for n > 0 {
		digit := n % 2
		res += strconv.Itoa(digit)
		n /= 2
	}
	return res
}
