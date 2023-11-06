package leetcode 

func differenceOfSum(nums []int) int {
    elementSum := 0
    digitSum := 0
    
    for i := 0; i < len(nums); i++ {
        digitSum += nums[i]
        elementSum += parseDigit(nums[i])
    }
    
    if digitSum - elementSum < 0 {
        return elementSum - digitSum
    }
    
    return digitSum - elementSum
}

func parseDigit(n int) int {
    if n < 10 {
        return n
    }
    
    sum := 0
    for n > 9 {
        sum += n % 10
        n /= 10
    }
    
    return sum + n
}
