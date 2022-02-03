package algo

import (
	// "fmt"
	. "../Tree"
)
func max(a int, b int)int{
	if a>=b{
		return a
	} else{
		return b
	}
}
func helper(root *TreeNode, res* int)int{
	if root == nil{
		return -2000
	}
	left := helper(root.Left, res)
	right := helper(root.Right, res)
	// fmt.Printf("node:%d,left:%d,right:%d\n",root.Val,left,right)
	sumNotRoot := max(max(left, right)+root.Val, root.Val)
	sumRoot := max(sumNotRoot, left+right+root.Val)
	*res = max(*res, sumRoot)
	return sumNotRoot
}
func maxPathSum(root *TreeNode) int {
	res := -3000
    helper(root, &res)
    return res
}