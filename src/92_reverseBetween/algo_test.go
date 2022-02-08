package algo

import (
	. "../ListNode"
	"fmt"
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	nodes := ArrToListNode([]int{1, 2, 3, 4, 5})
	expected := ArrToListNode([]int{1, 4, 3, 2, 5})
	actual := reverseBetween(nodes, 2, 4)
	fmt.Printf("actual:%v\n", ArrFromListNode(actual))
	if reflect.DeepEqual(actual, expected) {
		t.Log("Success")
	} else {
		t.Error("Fail")
	}

}
func TestReverseList2(t *testing.T) {
	nodes := ArrToListNode([]int{3, 5})
	expected := ArrToListNode([]int{5, 3})
	actual := reverseBetween(nodes, 1, 2)
	fmt.Printf("actual:%v\n", ArrFromListNode(actual))
	if reflect.DeepEqual(actual, expected) {
		t.Log("Success")
	} else {
		t.Error("Fail")
	}

}

func TestReverseList3(t *testing.T) {
	nodes := ArrToListNode([]int{1, 2, 3, 4})
	expected := ArrToListNode([]int{4, 3, 2, 1})
	actual := reverseBetween(nodes, 1, 4)
	fmt.Printf("actual:%v\n", ArrFromListNode(actual))
	if reflect.DeepEqual(actual, expected) {
		t.Log("Success")
	} else {
		t.Error("Fail")
	}

}
