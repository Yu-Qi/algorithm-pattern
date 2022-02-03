package algo
import (
	"container/list"
	// "fmt"
	. "../Tree"
)

func visit(node *TreeNode, queue *list.List, level_item *[]int){
	if node == nil{
		// fmt.Printf("node visit, position:%d, value:null\n",idx)
		return
	}
	*level_item = append(*level_item, node.Val)
	// fmt.Printf("node visit,  position:%d, value:%d\n",idx, node.Val)
	queue.PushBack(node.Left)
	queue.PushBack(node.Right)
}
func Pop(queue *list.List)*TreeNode{
	item := queue.Front()
	queue.Remove(queue.Front())
	// ***
	return item.Value.(*TreeNode)
}
func PushAll(queue *list.List, res *[][]int){
	size := queue.Len()
	level_item := []int{}
	for idx:= 0;idx<size;idx++{
		item := Pop(queue)
		visit(item,queue, &level_item)
	}
	if len(level_item) != 0{
		*res = append([][]int{level_item}, *res...)
	}
}

func levelOrderBottom(root *TreeNode)[][]int{
	if root == nil{
		return nil
	}
	queue := list.New()
	res := [][]int{}
	queue.PushBack(root)
	for queue.Len() > 0{
		PushAll(queue, &res)
	}

	return res
}