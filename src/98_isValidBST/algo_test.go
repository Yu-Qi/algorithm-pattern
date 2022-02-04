package algo
import (
	"testing"
	. "../covert_array_to_cbt"
)

func TestBalanced1(t *testing.T){
	nums := []int{5,1,4,0,0,3,6}
	root := CreateCompleteBibaryTree(nums)
	if isValidBST(root) ==false {
		t.Log("Success")
	} else{
		t.Error("Fail")
	}
}

func TestBalanced2(t *testing.T){
	nums := []int{2,2,2}
	root := CreateCompleteBibaryTree(nums)
	if isValidBST(root) ==false {
		t.Log("Success")
	} else{
		t.Error("Fail")
	}
}

func TestBalanced3(t *testing.T){
	nums := []int{5,4,6,0,0,3,7}
	root := CreateCompleteBibaryTree(nums)
	if isValidBST(root) ==false {
		t.Log("Success")
	} else{
		t.Error("Fail")
	}
}