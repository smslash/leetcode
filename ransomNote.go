package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[byte]int)

	for i := 0; i < len(ransomNote); i++ {
		m[ransomNote[i]]++
	}

	for i := 0; i < len(magazine); i++ {
		_, ok := m[magazine[i]]
		if ok {
			m[magazine[i]]--
		}
	}

	for _, v := range m {
		if v > 0 {
			return false
		}
	}


	return true
}
