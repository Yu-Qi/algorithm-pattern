package algo
import (
	. "../ListNode"
	"reflect"
	"testing"
)

func TestSort(t *testing.T){
	node1 := ArrToListNode([]int{4,2,1,3})
	actual := sortList(node1)
	expected := ArrToListNode([]int{1,2,3,4})
	if reflect.DeepEqual(actual, expected){
		t.Log("Success")
	} else{
		t.Error("Fail")
	}
}


func TestSort2(t *testing.T){
	node1 := ArrToListNode([]int{-1,5,3,4,0})
	actual := sortList(node1)
	expected := ArrToListNode([]int{-1,0,3,4,5})
	if reflect.DeepEqual(actual, expected){
		t.Log("Success")
	} else{
		t.Error("Fail")
	}

}

func TestSort3(t *testing.T){
	node1 := ArrToListNode([]int{})
	actual := sortList(node1)
	expected := ArrToListNode([]int{})
	if reflect.DeepEqual(actual, expected){
		t.Log("Success")
	} else{
		t.Error("Fail")
	}

}