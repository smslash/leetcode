package leetcode

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    var leaves1, leaves2 []int
    collectLeaves(root1, &leaves1)
    collectLeaves(root2, &leaves2)
    
    if len(leaves1) != len(leaves2) {
        return false
    }
    for i := range leaves1 {
        if leaves1[i] != leaves2[i] {
            return false
        }
    }
    return true
}

func collectLeaves(node *TreeNode, leaves *[]int) {
    if node == nil {
        return
    }
    if node.Left == nil && node.Right == nil {
        *leaves = append(*leaves, node.Val)
        return
    }
    collectLeaves(node.Left, leaves)
    collectLeaves(node.Right, leaves)
}
