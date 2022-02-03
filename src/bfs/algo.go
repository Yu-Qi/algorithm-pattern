package bfs

import (
	"container/list"
	"fmt"
	. "../Tree"
)

func visit(node *TreeNode, queue *list.List, idx int){
	if node == nil{
		fmt.Printf("node visit, position:%d, value:null\n",idx)
		return
	}
	fmt.Printf("node visit,  position:%d, value:%d\n",idx, node.Val)
	queue.PushBack(node.Left)
	queue.PushBack(node.Right)
}
func Pop(queue *list.List)*TreeNode{
	item := queue.Front()
	queue.Remove(queue.Front())
	// ***
	return item.Value.(*TreeNode)
}
func PushAll(queue *list.List){
	size := queue.Len()
	for idx:= 0;idx<size;idx++{
		item := Pop(queue)
		visit(item,queue, idx)
	}
}
func BFS(root *TreeNode){
	if root == nil{
		return
	}
	level := 1
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0{
		fmt.Printf("~~~ level: %d ~~~\n",level)
		PushAll(queue)
		level ++
	}
}