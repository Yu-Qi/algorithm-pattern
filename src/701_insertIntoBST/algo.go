package algo
import (
	. "../Tree"
)
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil{
        return &TreeNode{val,nil,nil}
    }
	insert(root, val)
	return root
}
func insert(root *TreeNode, val int){
	if root.Val >= val{
        if root.Left == nil{
            root.Left = &TreeNode{val, nil, nil}                    
        }else{
    		insertIntoBST(root.Left, val)
        }
	} else{
        if root.Right == nil{
            root.Right = &TreeNode{val,nil,nil}
        } else{
		insertIntoBST(root.Right, val)
        }
	}
}