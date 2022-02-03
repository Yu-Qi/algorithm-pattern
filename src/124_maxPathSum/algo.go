package algo

import (
	"fmt"
	. "../Tree"
)

func maxPathSum(root *TreeNode) int {
	if root == nil{
		return -2000
	}
	sum_left := maxPathSum(root.Left)
	sum_right := maxPathSum(root.Right)
	fmt.Printf("node:%d,left:%d,right:%d\n",root.Val,sum_left,sum_right)
	sum_total := root.Val
	if sum_left >0{
		sum_total += sum_left
	}
	if sum_right >0 {
		sum_total += sum_right
	}
	if sum_total >= sum_left && sum_total >= sum_right{
		return sum_total
	} else if sum_left > sum_right{
		return sum_left
	} else{
		return sum_right
	}
}