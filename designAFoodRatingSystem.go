package leetcode

type food struct {
	name    string
	cuisine string
	rating  int

	heapIdx int
}

type foodHeap []*food

func (rv *foodHeap) Len() int {
	return len(*rv)
}

func (rv *foodHeap) Less(i, j int) bool {
	if (*rv)[i].rating == (*rv)[j].rating {
		return (*rv)[i].name < (*rv)[j].name
	}
	return (*rv)[i].rating > (*rv)[j].rating

}

func (rv *foodHeap) Swap(i, j int) {
	(*rv)[i], (*rv)[j] = (*rv)[j], (*rv)[i]
	(*rv)[i].heapIdx = i
	(*rv)[j].heapIdx = j
}

func (rv *foodHeap) Push(x any) {
	*rv = append(*rv, x.(*food))
}

func (rv *foodHeap) Pop() any {
	l := rv.Len()
	toRemove := (*rv)[l-1]

	*rv = (*rv)[:l-1]

	return toRemove
}

type FoodRatings struct {
	cuisines map[string]*foodHeap
	tracker  map[string]*food
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	cuisineHeaps := make(map[string]*foodHeap)
	tracker := make(map[string]*food)

	for i := 0; i < len(foods); i++ {
		if _, ok := cuisineHeaps[cuisines[i]]; !ok {
			cuisineHeaps[cuisines[i]] = &foodHeap{}
		}

		f := &food{
			name:    foods[i],
			cuisine: cuisines[i],
			rating:  ratings[i],
			heapIdx: len(*cuisineHeaps[cuisines[i]]),
		}

		heap.Push(cuisineHeaps[cuisines[i]], f)

		tracker[foods[i]] = f
	}
	fr := FoodRatings{
		cuisines: cuisineHeaps,
		tracker:  tracker,
	}

	for _, h := range fr.cuisines {
		heap.Init(h)
	}

	return fr
}

func (rv *FoodRatings) ChangeRating(food string, newRating int) {
	f := rv.tracker[food]

	f.rating = newRating
	heap.Fix(rv.cuisines[f.cuisine], f.heapIdx)
}

func (rv *FoodRatings) HighestRated(cuisine string) string {
	return (*rv.cuisines[cuisine])[0].name
}
