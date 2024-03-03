package leetcode

func twoSum(numbers []int, target int) []int {
    i, j := 0, 1

    for j < len(numbers){
        if numbers[i] + numbers[j] == target {
            return []int{i+1,j+1}
        }else if target < numbers[i] + numbers[j] {
            i--
        }else{
            j++
            i++
        }   
    }
    return []int{i,j}
}
