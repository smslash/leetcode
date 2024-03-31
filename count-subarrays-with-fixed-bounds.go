package leetcode

func countSubarrays(nums []int, minK int, maxK int) int64 {
    if len(nums) == 0{
        return 0
    }

    minIndex, maxIndex := -1, -1
    res := 0
    left, right := 0, 0

    for right < len(nums){
        if nums[right] > maxK || nums[right] < minK{
            left = right+1
            minIndex, maxIndex = -1, -1
        }else{
            if nums[right]==minK{
                minIndex=right
            }

            if nums[right]==maxK{
                maxIndex=right
            }

            res += max(0,min(minIndex,maxIndex)-left+1)
        }
        right += 1
    }
    return int64(res)
}
