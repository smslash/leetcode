package leetcode

func findWinners(matches [][]int) [][]int {
    dp1:=make(map[int]int)
    dp2:=make(map[int]int)
    n:=len(matches)
    for i:=0; i<n; i++ {
        dp1[matches[i][0]]++
        dp2[matches[i][1]]++
    }
    var notLostAny []int
    var lostOnlyOne []int

    for player,_:=range dp1{
        if dp2[player]==0{
            notLostAny=append(notLostAny,player)
        }
    }
    for player,count:=range dp2{
        if count==1{
            lostOnlyOne=append(lostOnlyOne,player)
        }
    }
    sort.Ints(notLostAny)
    sort.Ints(lostOnlyOne)
    return [][]int{notLostAny,lostOnlyOne}
}
