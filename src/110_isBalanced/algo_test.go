package algo
import "testing"

// Input: root = [3,9,20,null,null,15,7]
func TestBalanced1(t *testing.T){
	// layer1
	l1 := TreeNode{3,nil,nil}
	// layer2
	l2_1 := TreeNode{9,nil,nil}
	l2_2 := TreeNode{20,nil,nil}
	// layer3
	l3_3 := TreeNode{15,nil,nil}
	l3_4 := TreeNode{7,nil,nil}

	l1.Left = &l2_1
	l1.Right = &l2_2
	l2_2.Left = &l3_3
	l2_2.Right = &l3_4
	if isBalanced(&l1) ==true{
		t.Log("Success")
	} else{
		t.Error("Fail")
	}

}


//Input: root = [1,2,2,3,3,null,null,4,4]
func TestBalanced2(t *testing.T){
	// layer1
	l1 := TreeNode{1,nil,nil}
	// layer2
	l2_1 := TreeNode{2,nil,nil}
	l2_2 := TreeNode{2,nil,nil}
	// layer3
	l3_1 := TreeNode{3,nil,nil}
	l3_2 := TreeNode{3,nil,nil}
	// layer4
	l4_1 := TreeNode{4,nil,nil}
	l4_2 := TreeNode{4,nil,nil}

	l1.Left = &l2_1
	l1.Right = &l2_2
	l2_1.Left = &l3_1
	l2_1.Right = &l3_2

	l3_1.Left = &l4_1
	l3_1.Right =&l4_2
	if isBalanced(&l1) ==false{
		t.Log("Success")
	} else{
		t.Error("Fail")
	}

}
