package leetcode

func minimumPushes(word string) int {
    var answer int
    wordLen := len(word) 
    for wordLen > 0 {
        answer += wordLen
        wordLen -= 8
    }
    return answer
}
