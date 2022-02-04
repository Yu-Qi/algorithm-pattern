package inorderTraversal
import (
	"testing"
	"reflect"
	. "../covert_array_to_cbt"
)

func TestInorderTraversal(t *testing.T){
	expected := []int{1,5,3,4,6}
	nums := []int{5,1,4,0,0,3,6}
	root := CreateCompleteBibaryTree(nums)	
	actual := InorderTraversal(root)
	if reflect.DeepEqual(actual, expected){
		t.Log("Success")
	} else{
		t.Error("Fail")
	}

}
