package algo
// import "fmt"

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	 }
func max(a int, b int)int{
	if a >= b{
		return a
	} else{
		return b
	}
}

func maxDepth(root *TreeNode) int {
	// fmt.Println(root)
	if root == nil || *root == (TreeNode{}){
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right))+1
}