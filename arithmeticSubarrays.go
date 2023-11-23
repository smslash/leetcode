package leetcode

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	answer := make([]bool, len(l))

	for i := 0; i < len(l); i++ {
		arr := make([]int, r[i]-l[i]+1)
		count := 0
		for x := l[i]; x <= r[i]; x++ {
			arr[count] = nums[x]
			count++
		}

		if isArithmetic(arr) {
			answer[i] = true
		}
	}
	return answer
}

func isArithmetic(a []int) bool {
	sort.Ints(a)
	diff := a[1] - a[0]
	for i := 1; i < len(a)-1; i++ {
		if a[i+1]-a[i] != diff {
			return false
		}
	}
	return true
}
