package leetcode

import "sort"

func minMovesToSeat(seats []int, students []int) int {
    sort.Ints(seats)
    sort.Ints(students)
    var ans int

    for i := 0; i < len(seats); i++ {
        if students[i] < seats[i] {
            ans += seats[i] - students[i]
        } else {
            ans += students[i] - seats[i]
        }
    }

    return ans
}
