package leetcode

func lengthOfLIS(nums []int) int {
    sub := []int{nums[0]}    

    for i := 1; i < len(nums); i++ {
        num := nums[i]
        if num > sub[len(sub)-1] {
            sub = append(sub, num)
        } else {
            j := sort.Search(len(sub), func(m int) bool {return sub[m] >= num})
            sub[j] = num
        }
    }
    
    return len(sub)
}
