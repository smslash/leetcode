package leetcode

func minOperations(nums []int) int {
    result := 0
    if len(nums)==0{
        return 0
    }
    numMap := make(map[int]int)

    for k :=0;k<len(nums);k++{
        numMap[nums[k]]++
    }

    for _,j := range numMap{
        if j%3==0{
            result = result + (j/3)
        }else if j%3==1 || j%3 == 2{
            if j == 1{
                return -1
            }
            result = result + (j/3) + 1
        }
    }

    return result
}
