package leetcode

func convert(s string, numRows int) string {
    zigzag := make([][]rune, numRows)
    
    i := 0
    col := 0
    for i < len(s) {
        for row := 0; row < numRows; row++ {
            zigzag[row] = append(zigzag[row], rune(s[i]))
            i++
            if i >= len(s) {
                break
            }
        }
        
        for row := numRows - 2; row > 0; row-- {
            if i >= len(s) {
                break
            }
            zigzag[row] = append(zigzag[row], rune(s[i]))
            col++
            i++
        }
    }
	
    zigzagStr := ""             
    for row := 0; row < numRows; row++ {
        zigzagStr += string(zigzag[row])    
    }
    return zigzagStr
}
