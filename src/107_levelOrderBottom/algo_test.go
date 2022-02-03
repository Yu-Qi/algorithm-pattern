package algo
import (
	"testing"
	. "../covert_array_to_cbt"
	"reflect"
)
func TestBalanced1(t *testing.T){
	nums := []int{3,9,20,0,0,15,7}
	root := CreateCompleteBibaryTree(nums)
	actual := levelOrderBottom(root)
	expect := [][]int{
		{15,7},
		{9,20},
		{3},
	}
	t.Log(actual)
	if reflect.DeepEqual(actual,expect) ==true {
		t.Log("Success")
	} else{
		t.Error("Fail")
	}
}