package algo
import (
	. "../Tree"
)

func inorderTraversal_recur(root *TreeNode) []int {
    if root == nil{
        return nil
    }
    left := inorderTraversal(root.Left)
    res := append(left, root.Val)
    right := inorderTraversal(root.Right)
    res = append(res, right...)
    return res
}

func inorderTraversal_v0(root *TreeNode) []int {
	stack := []*TreeNode{root}
	res := make([]int,0)

    for stack != nil && len(stack)>0{
        node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil{
			continue
		}
		if node.Left == nil && node.Right == nil{
			res = append(res,node.Val)
		} else{
            new_node :=  &TreeNode{node.Val,nil,nil}
            new_stack := []*TreeNode{node.Right,new_node,node.Left }
            stack = append(stack,new_stack...)
		}
	}
    return res
}

func inorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := make([]int,0)

    for root != nil || len(stack)>0{
		for root != nil{
			stack = append(stack,root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		root = node.Right
	}
		
    return res
}