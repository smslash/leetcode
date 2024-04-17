package leetcode

func smallestFromLeaf(root *TreeNode) string {
    var (
        result string
        visited = map[*TreeNode]bool{}
        s = []*TreeNode {root}
        path = []uint8 { uint8(root.Val) + 'a' }
        n *TreeNode
    )
    
	for len(s) > 0 {
		n = s[len(s)-1]
		visited[n] = true
		if n.Left == nil && n.Right == nil {
			s := string(path)
			if strings.Compare(s,result) == -1 || len(result) == 0 {
				result = s
			}
		}
		if n.Left != nil && !visited[n.Left]{
			s = append(s,n.Left)
			path = append([]uint8{ uint8(n.Left.Val) + 'a' }, path...)
			continue
		}
		if n.Right != nil && !visited[n.Right] {
			s = append(s,n.Right)
			path = append([]uint8{ uint8(n.Right.Val) + 'a' }, path...)
			continue
		}
        
		s = s[:len(s)-1]
		path = path[1:]
	}
	return result
}
