package Arr2Tree
import (
	"testing"
	. "../bfs"
)

//Input: root = [3,9,20,null,null,15,7]
func TestMaxDepth1(t *testing.T){
	nums := []int{3,9,20,0,0,15,7}
	root := CreateCompleteBibaryTree(nums)
	// ***
	BFS(root)
	// if maxDepth(&root) == 3{
	// 	t.Log("Success")
	// } else {
	// 	t.Error("Fail")
	// }
}

