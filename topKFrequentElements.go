package leetcode

func topKFrequent(nums []int, k int) []int {
	fq := make(map[int]int)
	a := []int{}

	for i := 0; i < len(nums); i++ {
		fq[nums[i]]++
	}

	tmp := make([]Num, len(fq))
	count := 0

	for i, v := range fq {
		tmp[count].i = i
		tmp[count].v = v
		count++
	}
	sortted(tmp)

	for i := 0; i < k; i++ {
		a = append(a, tmp[i].i)
	}

	return a
}

func sortted(a []Num) {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i].v < a[j].v {
				a[i].v, a[j].v = a[j].v, a[i].v
				a[i].i, a[j].i = a[j].i, a[i].i
			}
		}
	}
}

type Num struct {
	i int
	v int
}
