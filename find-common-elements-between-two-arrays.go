package leetcode

func findIntersectionValues(nums1 []int, nums2 []int) []int {
    mp1 := make(map[int]bool)
    mp2 := make(map[int]bool)
    
    for _, num := range nums1 {
        mp1[num] = true
    }
    
    for _, num := range nums2 {
        mp2[num] = true
    }
    
    c1, c2 := 0, 0
    
    for _, num := range nums1 {
        if mp2[num] {c1++}
    }

    for _, num := range nums2{
        if mp1[num] { c2++}
    }

    return []int{c1, c2}
}
