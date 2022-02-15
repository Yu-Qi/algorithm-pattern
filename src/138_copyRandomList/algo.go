package algo

type Node struct {
	     Val int
	     Next *Node
	     Random *Node
}/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */


func copyRandomList(head *Node) *Node {
    if head == nil{
        return nil
    }	
    // 1. Clone node
	cur := head
	for cur != nil{
		clone := &Node{Val:cur.Val, Next:cur.Next}
		cur.Next = clone
		cur = clone.Next
	}
	cur = head
	// 2. clone random
	for cur != nil{
        if cur.Random != nil { // random可能是nil
	    	cur.Next.Random = cur.Random.Next
	    }
        cur = cur.Next.Next
    }
	cur = head
	cloneHead := cur.Next
	// 3. split
	for cur != nil && cur.Next!=nil{
        tmp := cur.Next
		cur.Next = cur.Next.Next
		cur = tmp
	}
	return cloneHead
}