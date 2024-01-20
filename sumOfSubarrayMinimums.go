package leetcode

func sumSubarrayMins(arr []int) int {
	stack := []int{}
	sum := 0

	for right := 0; right <= len(arr); right++ {
		for len(stack) > 0 && (right == len(arr) || arr[stack[len(stack)-1]] >= arr[right]) {
			mini := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := -1
			if len(stack) > 0 {
				left = stack[len(stack)-1]
			}
            r := (mini - left) * (right - mini) % 1000000007
            s := (arr[mini] * r) % 1000000007
			sum += s
            sum %= 1000000007
		}

		stack = append(stack, right)
	}

	return sum
}
