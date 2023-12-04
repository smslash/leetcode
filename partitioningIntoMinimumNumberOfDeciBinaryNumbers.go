package leetcode

func minPartitions(n string) int {
	var a byte
    for i := 0; i < len(n); i++ {
        if a < n[i] {
			a = n[i]
		}
    }
	return Atoi(a)
}

func Atoi(n byte) int {
    switch n {
    case 48:
        return 0
    case 49:
        return 1
    case 50:
        return 2
    case 51:
        return 3
    case 52:
        return 4
    case 53:
        return 5
    case 54:
        return 6
    case 55:
        return 7
    case 56:
        return 8
    default:
        return 9
    }
}
