package bfs
import (
	"testing"
	. "../Tree"
)

func TestBFS(t *testing.T){
	root := TreeNode{3,nil,nil}
	// layer2
	l2_1 := TreeNode{9,nil,nil}
	l2_2 := TreeNode{20,nil,nil}
	// layer3
	l3_3 := TreeNode{15,nil,nil}
	l3_4 := TreeNode{7,nil,nil}

	root.Left = &l2_1
	root.Right = &l2_2

	l2_2.Left = &l3_3
	l2_2.Right = &l3_4	
	BFS(&root)
}
