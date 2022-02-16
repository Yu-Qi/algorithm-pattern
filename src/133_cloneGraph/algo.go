package algo
type Node struct {
	Val int
	Neighbors []*Node
}
   
func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	return clone(node, visited)
}
func clone(node *Node, visited map[*Node]*Node)*Node{
    if node==nil{
        return nil
    }

	if v,ok := visited[node];ok{
		return v
	} 
	clone_node := &Node{node.Val, make([]*Node, len(node.Neighbors))} 
	visited[node] = clone_node
	for i:=0;i<len(node.Neighbors);i++{
		clone_node.Neighbors[i] = clone(node.Neighbors[i], visited)
	}
	return clone_node
}