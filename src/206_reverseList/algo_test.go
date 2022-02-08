package algo

import (
	. "../ListNode"
	// "fmt"
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	nodes := ArrToListNode([]int{1, 2, 3, 4, 5})
	expected := ArrToListNode([]int{5, 4, 3, 2, 1})
	actual := reverseList(nodes)
	// fmt.Printf("actual:%v\n", ArrFromListNode(actual))
	if reflect.DeepEqual(actual, expected) {
		t.Log("Success")
	} else {
		t.Error("Fail")
	}

}

func TestReverseList2(t *testing.T) {
	nodes := ArrToListNode([]int{1})
	expected := ArrToListNode([]int{1})
	actual := reverseList(nodes)
	// fmt.Printf("actual:%v\n", ArrFromListNode(actual))
	if reflect.DeepEqual(actual, expected) {
		t.Log("Success")
	} else {
		t.Error("Fail")
	}

}

func TestReverseList3(t *testing.T) {
	nodes := ArrToListNode([]int{})
	expected := ArrToListNode([]int{})
	actual := reverseList(nodes)
	// fmt.Printf("actual:%v\n", ArrFromListNode(actual))
	if reflect.DeepEqual(actual, expected) {
		t.Log("Success")
	} else {
		t.Error("Fail")
	}

}
