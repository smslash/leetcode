package leetcode

func largestWordCount(messages []string, senders []string) string {
	s := make(map[string]int)
	for i := 0; i < len(senders); i++ {
		s[senders[i]] += len(strings.Fields(messages[i]))
	}

	max := 0
	for _, v := range s {
		if v > max {
			max = v
		}
	}

	ans := []string{}
	for i, v := range s {
		if v == max {
			max = v
			ans = append(ans, i)
		}
	}

	if len(ans) == 1 {
		return ans[0]
	}

	res := ""
	for i := 0; i < len(ans); i++ {
		if ans[i] > res {
			res = ans[i]
		}
	}

	return res
}
