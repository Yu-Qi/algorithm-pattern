package inorderTraversal

import (
	"fmt"
	. "../Tree"
)
func Travesal(root *TreeNode)[]int{
	if root == nil{
		return nil
	}

	left := Travesal(root.Left)
	items := append(left, root.Val)
	right := Travesal(root.Right)
	items = append(items, right...)
	return items
}

func InorderTraversal(root *TreeNode) []int {
	res := Travesal(root)
	fmt.Println(res)
	return res
}