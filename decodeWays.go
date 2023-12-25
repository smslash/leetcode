package leetcode

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}

	tmp := make([]int, len(s)+1)
	tmp[0] = 1

	return recursive(tmp, s)
}

func recursive(arr []int, s string) int {
	if arr[len(s)] > 0 {
		return arr[len(s)]
	}

	if s[0] == '0' {
		return 0
	}

	sum := recursive(arr, s[1:])
	if len(s) >= 2 && s[:2] <= "26" {
		sum += recursive(arr, s[2:])
	}

	arr[len(s)] = sum

	return arr[len(s)]
}
