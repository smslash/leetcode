package leetcode

func countDistinctIntegers(nums []int) int {
	m := make(map[int]int)

	a := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		str := strconv.Itoa(nums[i])
		str = reverse(str)
		n, _ := strconv.Atoi(str)
		a[i] = n
	}

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		m[a[i]]++
	}

	return len(m)
}

func reverse(s string) string {
	a := make([]byte, len(s))
	j := 0

	for i := len(s) - 1; i >= 0; i-- {
		a[j] = s[i]
		j++
	}

	return string(a)
}
