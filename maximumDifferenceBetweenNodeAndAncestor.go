package leetcode

func maxAncestorDiff(root *TreeNode) int {
	var dfs func(*TreeNode, int, int) int
	dfs = func(node *TreeNode, minVal, maxVal int) int {
		if node == nil {
			return maxVal - minVal
		}
		minVal = min(minVal, node.Val)
		maxVal = max(maxVal, node.Val)
		leftDiff := dfs(node.Left, minVal, maxVal)
		rightDiff := dfs(node.Right, minVal, maxVal)
		return max(leftDiff, rightDiff)
	}
	return dfs(root, root.Val, root.Val)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
