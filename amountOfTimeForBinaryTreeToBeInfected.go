package leetcode

func amountOfTime(root *TreeNode, start int) int {
    adjacencyList := make(map[int][]int, 0)
    
    stack := []*TreeNode{root}
    
    for len(stack) > 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        
        if node == nil {
            continue
        }
        if node.Left != nil {
            adjacencyList[node.Val] = append(adjacencyList[node.Val], node.Left.Val)
            adjacencyList[node.Left.Val] = append(adjacencyList[node.Left.Val], node.Val)
        }
        
        if node.Right != nil {
            adjacencyList[node.Val] = append(adjacencyList[node.Val], node.Right.Val)
            adjacencyList[node.Right.Val] = append(adjacencyList[node.Right.Val], node.Val)
        }
        
        stack = append(stack, node.Left, node.Right)
    }
    
    queue := [][2]int{{start, 0}}
    visited := make(map[int]bool)
    result := 0
    
    for len(queue) > 0 {
        node, level := queue[0][0], queue[0][1]
        
        queue = queue[1:]
        
        if visited[node] {
            continue
        }
        visited[node] = true
        
        if level > result {
            result = level
        }
        
        for _, childNode := range adjacencyList[node] {
            queue = append(queue, [2]int{childNode, level+1})
        }
    }
    
    return result
}
