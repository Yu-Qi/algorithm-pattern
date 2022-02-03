package Arr2Tree
import (
	"fmt"
	. "../Tree"
)

func AddNode(node *TreeNode, nums []int, idx int, size int)*TreeNode{
	// ***
	if idx > size || nums[idx-1] == 0{
		return nil
	}
	node.Val = nums[idx-1]
	node.Left = AddNode(&TreeNode{}, nums, idx*2, size)
	node.Right = AddNode(&TreeNode{}, nums, idx*2+1, size)
	return node
}

func CreateCompleteBibaryTree(nums []int)*TreeNode{
	if nums == nil{
		return nil
	}

	size := len(nums)
	root := TreeNode{}
	AddNode(&root, nums, 1, size)
	return &root
}

func TravesalTree(node *TreeNode){
	if node == nil{
		return 
	}
	fmt.Println(node.Val)
	
}