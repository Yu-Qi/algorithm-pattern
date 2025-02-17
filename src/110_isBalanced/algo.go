package algo
import (
	"fmt"
	. "../Tree"
)

func max(a int, b int)int{
	if a>=b{
		return a
	} else{
		return b
	}
}
func Abs(num int)int{
	if num >=0{
		return num
	} else{
		return -1*num
	}
}
func GetDepth(root *TreeNode)(int,bool){
	if root == nil{
		return 0,true
	}
	depth_left, balanced_left :=  GetDepth(root.Left)
	depth_right, balanced_right :=  GetDepth(root.Right)

	fmt.Printf("value=%d,depth_left=%d,balanced_left=%t,depth_right=%d,balanced_right=%t\n",root.Val,depth_left,balanced_left,depth_right,balanced_right)
	if !balanced_left|| !balanced_right{
		return -1,false
	} else if Abs(depth_left-depth_right) >=2 {
		return -1,false
	} else{
		return max(depth_left+1,depth_right+1),true
	}
}

func isBalanced(root *TreeNode) bool {
	_, ok := GetDepth(root)
	return ok
}