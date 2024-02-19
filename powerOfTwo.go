package leetcode

func isPowerOfTwo(n int) bool {
    if n == 1 {
        return true
    }
    
    if n % 2 != 0 || n == 0 {
        return false
    }

    for {
        if n == 1 {
            return true
        }

        if n % 2 != 0 {
            break
        }

        n /= 2
    }

    return false
}
