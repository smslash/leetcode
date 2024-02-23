package leetcode

type Pair struct {
	to, cost int
}

type Data struct {
	cost, next int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	graph := make([][]Pair, n)
	for _, conn := range flights {
		from := conn[0]
		to := conn[1]
		cost := conn[2]
		graph[from] = append(graph[from], Pair{to: to, cost: cost})
	}
	queue := make([]Data, 0)
	distance := make([]int, n)
	for i := 0; i < n; i++ {
		distance[i] = math.MaxInt16
	}
	stops := k + 1
	queue = append(queue, Data{cost: 0, next: src})
	for len(queue) > 0 && stops > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			it := queue[0]
			queue = queue[1:]
			currCost, currPlace := it.cost, it.next
			for _, val := range graph[currPlace] {
				toCost, toPlace := val.cost, val.to
				if currCost+toCost < distance[toPlace] {
					distance[toPlace] = currCost + toCost
					queue = append(queue, Data{cost: distance[val.to], next: val.to})
				}
			}
		}
		stops--
	}
	if distance[dst] == math.MaxInt16 {
		return -1
	}
	return distance[dst]
}
