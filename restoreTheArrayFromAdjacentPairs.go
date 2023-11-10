package leetcode

func restoreArray(adjacentPairs [][]int) []int {
    if len(adjacentPairs) == 0 {
        return []int{}
    }

    adj := make(map[int][]int)
    for _, pair := range adjacentPairs {
        adj[pair[0]] = append(adj[pair[0]], pair[1])
        adj[pair[1]] = append(adj[pair[1]], pair[0])
    }

    start := 0
    for key, val := range adj {
        if len(val) == 1 {
            start = key
            break
        }
    }

    result := make([]int, 0, len(adjacentPairs)+1)
    result = append(result, start)
    prev := start
    for len(result) < len(adjacentPairs)+1 {
        next := adj[prev][0]
        if len(result) > 1 && next == result[len(result)-2] {
            next = adj[prev][1]
        }
        result = append(result, next)
        prev = next
    }

    return result
}
