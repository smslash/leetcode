package leetcode

func findGCD(nums []int) int {
	sort.Ints(nums)
	min := 0

	if isPrime(nums[0]) {
		min = nums[0]
	} else {
		min = 2
	}

	for i := nums[len(nums)-1]; i >= min; i-- {
		if nums[0]%i == 0 && nums[len(nums)-1]%i == 0 {
			return i
		}
	}

	return 1
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
