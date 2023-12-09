package leetcode

func inorderTraversal(root *TreeNode) []int {
  result := make([]int, 0)
	if root != nil {
		if root.Left != nil {
			result = append(result, inorderTraversal(root.Left)...)
		}
		result = append(result, root.Val)
		if root.Right != nil {
			result = append(result, inorderTraversal(root.Right)...)
		}
	}
	return result
}
