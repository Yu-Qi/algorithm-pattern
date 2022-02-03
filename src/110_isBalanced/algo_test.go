package algo
import (
	"testing"
	. "../covert_array_to_cbt"
)
// Input: root = [3,9,20,null,null,15,7]
func TestBalanced1(t *testing.T){
	nums := []int{3,9,20,0,0,15,7}
	root := CreateCompleteBibaryTree(nums)
	if isBalanced(root) ==true{
		t.Log("Success")
	} else{
		t.Error("Fail")
	}

}