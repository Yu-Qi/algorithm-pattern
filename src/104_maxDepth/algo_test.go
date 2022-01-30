package algo
import "testing"

//Input: root = [3,9,20,null,null,15,7]
func TestMaxDepth1(t *testing.T){
	// layer1
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
	if maxDepth(&root) == 3{
		t.Log("Success")
	} else {
		t.Error("Fail")
	}
}

//Input: root = []
func TestMaxDepth2(t *testing.T){
	// layer1
	var root TreeNode
	if maxDepth(&root) == 0{
		t.Log("Success")
	} else {
		t.Errorf("Fail, answer=%d", maxDepth(&root))
	}
}
