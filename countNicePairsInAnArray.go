package leetcode 

func rev(num int) int {
	result := 0
	for num > 0 {
		result = result*10 + num%10
		num /= 10
	}
	return result
}

func countNicePairs(nums []int) int {
	arr := make([]int, len(nums))
	dic := make(map[int]int)
	ans := 0
	const MOD = 1000000007

	for i, num := range nums {
		arr[i] = num - rev(num)
	}

	for _, num := range arr {
		ans = (ans + dic[num]) % MOD
		dic[num]++
	}

	return ans
}
