package leetcode

func sortVowels(s string) string {
    vowels := []rune{}
    for _, ch := range s {
        if strings.ContainsRune("aeiouAEIOU", ch) {
            vowels = append(vowels, ch)
        }
    }

    sort.Slice(vowels, func(i, j int) bool {
        return vowels[i] < vowels[j]
    })

    vowelIndex := 0
    result := make([]rune, len(s))

    for i, ch := range s {
        if strings.ContainsRune("aeiouAEIOU", ch) {
            result[i] = vowels[vowelIndex]
            vowelIndex++
        } else {
            result[i] = ch
        }
    }

    return string(result)
}
