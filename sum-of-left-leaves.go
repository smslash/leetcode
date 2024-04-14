package leetcode

func sumOfLeftLeaves(root *TreeNode) int {
    if root == nil {
        return 0
    }

    var sum int
    if root.Left != nil {
        if root.Left.Left == nil && root.Left.Right == nil {
            sum += root.Left.Val
        } else {
            sum += sumOfLeftLeaves(root.Left) 
        }
    }
    
    sum += sumOfLeftLeaves(root.Right)
    
    return sum
}
