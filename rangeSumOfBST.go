package leetcode 

func rangeSumBSTUtil(root *TreeNode, low, high int, sum *int) {
	if root == nil {
		return
	}

	if low <= root.Val && root.Val <= high {
		*sum += root.Val
	}

	rangeSumBSTUtil(root.Left, low, high, sum)
	rangeSumBSTUtil(root.Right, low, high, sum)
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0
	rangeSumBSTUtil(root, low, high, &sum)
	return sum
}
