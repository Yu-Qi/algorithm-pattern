package algo
import (
	. "../ListNode"
	"reflect"
	"testing"
	"fmt"
)

func TestPartition(t *testing.T){
	node1 := ArrToListNode([]int{1,4,3,2,5,2})
	actual := partition(node1,3)

	node3 := ArrFromListNode(actual)
	fmt.Printf("%v\n",node3)
	expected := ArrToListNode([]int{1,2,2,4,3,5})
	if reflect.DeepEqual(actual, expected){
		t.Log("Success")
	} else{
		t.Error("Fail")
	}
}

func TestPartition2(t *testing.T){
	node1 := ArrToListNode([]int{1})
	actual := partition(node1,0)

	node3 := ArrFromListNode(actual)
	fmt.Printf("%v\n",node3)
	expected := ArrToListNode([]int{1})
	if reflect.DeepEqual(actual, expected){
		t.Log("Success")
	} else{
		t.Error("Fail")
	}
}
