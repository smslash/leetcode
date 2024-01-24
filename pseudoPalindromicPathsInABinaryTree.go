package leetcode

func pseudoPalindromicPaths(root *TreeNode) int {
    return dfs(root, [9]int{})
}

func dfs(root *TreeNode, digits [9]int) int {
    digits[root.Val-1]++

    switch {
        case root.Left == nil && root.Right == nil:
            return canBuildPalindromic(digits)
        case root.Left == nil:
            return  dfs(root.Right, digits)
        case root.Right == nil:
            return dfs(root.Left, digits)
        default:
            return dfs(root.Left, digits) + dfs(root.Right, digits)
    }
}

func canBuildPalindromic(nums [9]int) int {
    count := 0
    oddUse := false
    for _, num := range nums {
        count += num
        if num % 2 != 0 {
            if oddUse {
                return 0
            }
            oddUse = true
        }
    }
    if count == 0 {
        return 0
    }
    return 1
}
