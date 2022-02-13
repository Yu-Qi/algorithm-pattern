package algo
import (
	. "../ListNode"
	"reflect"
	"testing"
	"fmt"
)

func TestMergeTwoLists(t *testing.T){
	node1 := ArrToListNode([]int{1, 2, 4})
	node2 := ArrToListNode([]int{1, 3, 4})
	actual := mergeTwoLists(node1,node2)

	node3 := ArrFromListNode(actual)
	fmt.Printf("%v\n",node3)
	expected := ArrToListNode([]int{1, 1, 2, 3, 4, 4})
	if reflect.DeepEqual(actual, expected){
		t.Log("Success")
	} else{
		t.Error("Fail")
	}
}


func TestMergeTwoLists2(t *testing.T){
	actual := mergeTwoLists(nil,nil)
	if reflect.DeepEqual(actual, nil){
		t.Log("Success")
	} else{
		t.Error("Fail")
	}
}