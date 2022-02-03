package algo
import (
	"testing"
	. "../covert_array_to_cbt"
)

func TestMaxPathSum(t *testing.T){
	nums := []int{-10,9,20,0,0,15,7}
	root := CreateCompleteBibaryTree(nums)
	// t.Log(maxPathSum(root))
	if maxPathSum(root) ==42{
		t.Log("Success")
	} else{
		t.Error("Fail")
	}

}


func TestMaxPathSum2(t *testing.T){
	nums := []int{1,2,3}
	root := CreateCompleteBibaryTree(nums)
	// t.Log(maxPathSum(root))
	if maxPathSum(root) ==6{
		t.Log("Success")
	} else{
		t.Error("Fail")
	}

}

func TestMaxPathSum3(t *testing.T){
	nums := []int{-1,-2,-3}
	root := CreateCompleteBibaryTree(nums)
	// t.Log(maxPathSum(root))
	if maxPathSum(root) ==-1{
		t.Log("Success")
	} else{
		t.Error("Fail")
	}

}

func TestMaxPathSum4(t *testing.T){
	nums := []int{1,-2,-3,1,3,-2,0,-1}
	root := CreateCompleteBibaryTree(nums)
	// t.Log(maxPathSum(root))
	if maxPathSum(root) ==3{
		t.Log("Success")
	} else{
		t.Error("Fail")
	}

}