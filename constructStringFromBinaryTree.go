package leetcode

func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}
	res := fmt.Sprintf("%d", root.Val)
	if root.Left != nil {
		res += fmt.Sprintf("(%s)", tree2str(root.Left))
	} else if root.Right != nil {
		res += "()"
	}
	if root.Right != nil {
		res += fmt.Sprintf("(%s)", tree2str(root.Right))
	}
	return res
}
