package leetcode

func interpret(command string) string {
	var res string
	l := len(command)

	for i := 0; i < l; i++ {
		if command[i] == '(' && i+1 < l {
			if command[i+1] == 'a' {
				res += "al"
				i += 3
			} else {
				res += string('o')
				i++
			}
		} else {
			res += string(command[i])
		}
	}

	return res
}
