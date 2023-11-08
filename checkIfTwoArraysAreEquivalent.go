package leetcode

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    var s1 strings.Builder
    var s2 strings.Builder
    
    for i := 0; i < len(word1); i++ {
        s1.WriteString(word1[i])
    }
    
    for i := 0; i < len(word2); i++ {
        s2.WriteString(word2[i])
    }
    
    return s1.String() == s2.String()
}
