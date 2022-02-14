package algo
import (
	. "../ListNode"
	"reflect"
	"testing"
	// "fmt"
)

func TestSort(t *testing.T){
	node1 := ArrToListNode([]int{1,2,3,4})
	reorderList(node1)
	expected := ArrToListNode([]int{1,4,2,3})
	if reflect.DeepEqual(node1, expected){
		t.Log("Success")
	} else{
		t.Error("Fail")
	}
}

func TestSort2(t *testing.T){
	node1 := ArrToListNode([]int{1,2})
	reorderList(node1)
	expected := ArrToListNode([]int{1,2})
	if reflect.DeepEqual(node1, expected){
		t.Log("Success")
	} else{
		t.Error("Fail")
	}
}

func TestSort3(t *testing.T){
	node1 := ArrToListNode([]int{1})
	reorderList(node1)
	expected := ArrToListNode([]int{1})
	if reflect.DeepEqual(node1, expected){
		t.Log("Success")
	} else{
		t.Error("Fail")
	}
}

func TestSort4(t *testing.T){
	node1 := ArrToListNode([]int{1,2,3,4,5})
	reorderList(node1)
	expected := ArrToListNode([]int{1,5,2,4,3})
	if reflect.DeepEqual(node1, expected){
		t.Log("Success")
	} else{
		t.Error("Fail")
	}
}
