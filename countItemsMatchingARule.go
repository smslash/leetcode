package leetcode

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	count := 0
	for i := 0; i < len(items); i++ {
		if ruleKey == "color" && string(items[i][1]) == ruleValue {
			count += 1
			continue
		} else if ruleKey == "type" && string(items[i][0]) == ruleValue {
			count += 1
			continue
		} else if ruleKey == "name" && string(items[i][2]) == ruleValue {
			count += 1
			continue
		}
	}
	return count
}
