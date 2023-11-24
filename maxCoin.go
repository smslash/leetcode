package leetcode

func maxCoins(piles []int) (res int) {
    sort.Ints(piles)
    for i:=len(piles) / 3; i < len(piles); i += 2 {
        res += piles[i]
    } 
    return res
}
