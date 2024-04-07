package leetcode

func groupThePeople(groupSizes []int) [][]int {  
    sizes := make(map[int][]int)
    for i, size := range groupSizes {
        sizes[size] = append(sizes[size], i)
    }

    ans := make([][]int, 0, len(groupSizes))
    for size, members := range sizes {
        for i := 0; i < len(members); i += size {
            ans = append(ans, members[i:i+size])
        }
    }

    return ans
}
