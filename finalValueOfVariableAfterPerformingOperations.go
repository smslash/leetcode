package leetcode

func finalValueAfterOperations(operations []string) int {
	res := 0
	for i := 0; i < len(operations); i++ {
		if operations[i] == "X--" || operations[i] == "--X" {
			res--
		} else if operations[i] == "X++" || operations[i] == "++X" {
			res++
		}
	}
	return res
}
