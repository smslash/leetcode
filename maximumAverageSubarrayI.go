package leetcode

func findMaxAverage(nums []int, k int) float64 {
    if len(nums) == k {
        sum := 0
        for _, num := range nums {
            sum += num
        }
        return float64(sum) / float64(k)
    }

    maxSum := 0
    currentSum := 0

    for i := 0; i < k; i++ {
        currentSum += nums[i]
    }
    maxSum = currentSum

    for i := k; i < len(nums); i++ {
        currentSum += nums[i] - nums[i-k]
        if currentSum > maxSum {
            maxSum = currentSum
        }
    }

    return float64(maxSum) / float64(k)
}
